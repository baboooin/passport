package main

import (
	"fmt"
	"log"
	"time"

	sms "github.com/baboooin/go_sms_ru"
	"github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)




func CheckUser(l LoginRec) (*UserRecord, error) {
	var user UserRecord
	var query string = "Select uid, firstname, lastname, hash from users where "
	db, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}

	//hash := bcrypt.GenerateFromPassword([]byte(password))
	if patternEmail.MatchString(l.Login) {
		query = query + "email='" + l.Login + "'"
	}

	if patternMobile.MatchString(l.Login) {
		query = query + "mobile='" + l.Login + "'"
	}

	log.Println(query)
	err = db.QueryRow(query).Scan(
		&user.UId,
		&user.FistsName,
		&user.LastName,
		&user.Hash,
	)
	db.Close()
	if err != nil {
		err = errors.New("Error login or password")
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(l.Password))
	if err != nil {
		// err = errors.New("Error make token")
		return nil, err
	}
	user.Hash = ""
	return &user, err
}

// func getUser(login string, p)

func CreateUser(l RegisterRec) (UserRecord, error) {
	var u UserRecord
	password, err := bcrypt.GenerateFromPassword([]byte(l.Password), 10)
	if err != nil {
		return u, err
	}

	db, err := ConnectDatabase()
	if err != nil {
		return u, err
	}

	count := 0
	if l.Email != "" {
		err = db.QueryRow("Select count(uid) from users where email=$1", l.Email).Scan(&count)
		if err != nil {
			return u, err
		}

		if count != 0 {
			return u, errors.New("User already registered!")
		}
		// log.Println(count)
	}

	if l.Mobile != "" {
		err = db.QueryRow("Select count(uid) from users where mobile=$1", l.Mobile).Scan(&count)
		if err != nil {
			return u, err
		}

		if count != 0 {
			return u, errors.New("User already registered!")
		}

		// log.Println(count)

	}

	err = db.QueryRow("INSERT INTO users(firstname, lastname, email, mobile, hash ) VALUES ($1, $2, $3, $4, $5) RETURNING uid, firstname, lastname, email, mobile", l.FistsName, l.LastName, l.Email, l.Mobile, string(password)).Scan(&u.UId, &u.FistsName, &u.LastName, &u.Email, &u.Mobile)
	if err != nil {
		return u, err
	}
	return u, nil
}

func UserExist(login string) (id string, ok bool) {
	// log.Println(" isEmail", patternEmail.MatchString(login))
	// log.Println(" isMobile", patternMobile.MatchString(login))

	if !patternMobile.MatchString(login) && !patternEmail.MatchString(login) {
		log.Println("Pattern error ")
		return "", false
	}

	var query string = "Select uid from users where "

	if patternEmail.MatchString(login) {
		query = query + "email like ('" + login + "')"
	}

	if patternMobile.MatchString(login) {
		query = query + "mobile like ('" + login + "')"
	}

	db, err := ConnectDatabase()
	if err != nil {
		return "", err != nil
	}

	err = db.QueryRow(query).Scan(&id)
	db.Close()
	return id, err == nil
}

func OTPLastupdate(uid string, minutes int) (left time.Duration, limit bool, ok bool) {
	db, err := ConnectDatabase()
	var t1 = time.Time{}
	if err != nil {
		return 0, false, false
	}

	var query string = "Select generationdate from otp where uid = $1;"
	err = db.QueryRow(query, uid).Scan(&t1)

	fiveMin := time.Minute * time.Duration(minutes)
	dateDeadLine := t1.Add(fiveMin)

	log.Printf("\ntime now: %v \ntime last call: %v\n five minutes: %v \n Until: %v", time.Now(), dateDeadLine, time.Until(dateDeadLine))

	return time.Until(dateDeadLine), time.Now().After(dateDeadLine), err == nil
}

func OTPUpdate(id string, code string) error {
	db, err := ConnectDatabase()
	if err != nil {
		return err
	}
	var query string = "update otp set code=$1, generationdate=$2 where uid=$3"
	log.Println(code, time.Now(), id)
	_, err = db.Exec(query, code, time.Now(), id) // .QueryRow(, uid, code).Scan(&generationdate)
	return err

}

func OTPInsert(uid string, code string) error {
	var generationdate time.Time
	db, err := ConnectDatabase()
	if err != nil {
		return err
	}
	var query string = "Insert into otp (uid, code) values ($1,$2) RETURNING generationdate"
	err = db.QueryRow(query, uid, code).Scan(&generationdate)
	log.Println(generationdate)
	return err
}

func CallSMSRU(Mobile string) (string, error) {
	client := sms.SmsRU(SMSRU_API_ID)
	call, err := client.Call(Mobile)
	if err != nil {
		return "", err
	}
	if call["status"].(string) != "OK" {
		return "", errors.New("SMSRU Call failed")
	}

	return fmt.Sprintf("%f", call["code"].(float64)), nil
}

func CallSMSRUfake(Mobile string) (string, error) {
	return GenCode(), nil
}
