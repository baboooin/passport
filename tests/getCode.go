package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"

	"github.com/pkg/errors"
)

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
