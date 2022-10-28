package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func GenRandomPassword() string {
	var (
		lowerCharSet   = "abcdedfghijklmnopqrst"
		upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		specialCharSet = "!@#$%&*"
		numberSet      = "0123456789"
		allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
	)

	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 8

	var password strings.Builder
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}


func GenCode() string {
	numberSet := "0123456789"
	rand.Seed(time.Now().Unix())
	minNum := 4
	var password strings.Builder
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}


func SendPassword(l RegisterRec) {

	// c := sms.SmsRU("26A73D0D-5116-8AF6-B031-2E9E6CD4483F")

	// if res, err := c.Call(l.Mobile); err == nil {
	// 	log.Println(res)
	// }

	// if res, err := c.Balance(); err == nil {
	// 	log.Println(res)
	// }

	// if res, err := c.Limit(); err == nil {
	// 	log.Println(res)
	// }

	// if res, err := c.Send("+ Mobile number", "Message"); err == nil {
	// 	log.Println(res)
	// }

	// if res, err := c.Call("+ Mobile number"); err == nil {
	// 	log.Println(res)
	// }

	// c.
	if l.Mobile != "" { // try send password to mobile

		// c := sms.SmsRU("26A73D0D-5116-8AF6-B031-2E9E6CD4483F")

		// param1:= &openapi.CreateMessageParams{}

		// client := twilio.NewRestClient()
		// message, err := client.Api.CreateMessage()

		// 	params := &openapi.CreateMessageParams{}
		// 	params.SetTo(l.Mobile)
		// 	params.SetFrom("+15017250604")
		// 	params.SetBody("Password: "+l.Password)

		// 	resp, err := client.LookupsV1.FetchPhoneNumber(l.Mobile,verifyParam) // .ApiV2010.CreateMessage(params)
		// 	if err != nil {
		// 		fmt.Println(err.Error())
		// 		err = nil
		// 	} else {
		// 		fmt.Println("Message Sid: " + *resp.Sid)
		// 	}

		// } else { // send to email
		// 	err := SendMailTempotaryPassword(
		// 		l.Email,
		// 		l.FistsName+" "+l.LastName,
		// 		"Register new user",
		// 		"templates/register_email_template.en-EN",
		// 		l.Password)

		// 	// err = SendMail(l.Email, body)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return
		// 	}
	}
}


func getCode(w http.ResponseWriter, r *http.Request) {

	var l MobileLoginRec

	//get
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		Error(err, w)
		return
	}
	log.Println(string(bodyBytes))

	if err := json.Unmarshal(bodyBytes, &l); err != nil {
		log.Println("Decode body ", err)
		Error(err, w)
		return
	}

	if !patternMobile.MatchString(l.Mobile) {
		Error(errors.New("Mobile error"), w)
		return
	}

	id, ok := UserExist(l.Mobile)
	if ok {
		if d, x, ok := OTPLastupdate(id, 5); ok {
			// log.Println("limit ", x)
			if !x {
				Error(errors.New(fmt.Sprintf("Сall limit. Try after %v minutes (%v)", math.Ceil(d.Minutes()), d.Minutes())), w)
				return
			}
			// log.Println("Update code")
			call, err := CallSMSRU(l.Mobile)
			if err != nil {
				Error(err, w)
				return
			}
			err = OTPUpdate(id, call)
			if err != nil {
				Error(err, w)
				return
			}

		} else {
			// insert code
			// log.Println("Insert code")
			call, err := CallSMSRU(l.Mobile)
			if err != nil {
				Error(err, w)
				return
			}

			err = OTPInsert(id, call)
			if err != nil {
				Error(err, w)
				return
			}
		}
	} else {
		// log.Println("User not found")
		// insert user
		var rr RegisterRec
		var ur UserRecord
		rr.Mobile = l.Mobile
		ur, err := CreateUser(rr)
		if err != nil {
			Error(err, w)
			return
		}
		// register  new code
		call, err := CallSMSRU(l.Mobile)
		if err != nil {
			Error(err, w)
			return
		}
		err = OTPInsert(ur.UId, call)
		if err != nil {
			Error(err, w)
			return
		}
	}

}