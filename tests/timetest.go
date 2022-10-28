package main

import (
	"errors"
	"fmt"
	"log"
	"math"

	sms "github.com/baboooin/go_sms_ru"
)

func CallSMSRU1(Mobile string) (string, error) {
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

func CallSMSRU(Mobile string) (string, error) {
	return GenCode(), nil
}

func main() {

	// var patternMobile1, _ = regexp.Compile(`^\+|\d[0-9]{7,12}$`) //^\+|\d[0-9]{7,12}$

	// // var patternEmail, _ = regexp.Compile(`[a-zA-Z0-9\\+\\.\\_\\%\\-\\+]{1,256}\\@[a-zA-Z0-9][a-zA-Z0-9\\-]{0,64}(\\.[a-zA-Z0-9][a-zA-Z0-9\\-]{0,25})+`)

	// // var patternEmail1, _ = regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,9}$`)
	var Mobile = "+79785507887"

	
	id, ok := UserExist(Mobile)
	if ok {
		if d, x, ok := OTPLastupdate(id, 5); ok {
			log.Println("limit ", x)
			if !x {
				log.Panicf("Сall limit. Try after %v minutes (%v)", math.Ceil(d.Minutes()), d.Minutes())
				return
			}
			log.Println("Update code")
			call, err := CallSMSRU(Mobile)
			if err != nil {
				log.Panic(err)
				return
			}
			err = OTPUpdate(id, call)

		} else {
			// insert code
			log.Println("Insert code")
			call, err := CallSMSRU(Mobile)
			if err != nil {
				log.Panic(err)
				return
			}

			err = OTPInsert(id, call)
			if err != nil {
				log.Panic(err)
				return
			}
		}
	} else {
		log.Println("User not found")
		// insert user
		var rr RegisterRec
		var ur UserRecord
		rr.Mobile = Mobile
		ur, err := CreateUser(rr)
		if err != nil {
			log.Panic(err)
		}
		// register  new code
		call, err := CallSMSRU(Mobile)
		if err != nil {
			log.Panic(err)
			return
		}
		err = OTPInsert(ur.UId, call)
		if err != nil {
			log.Panic(err)
			return
		}
	}
}
