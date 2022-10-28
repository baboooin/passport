package main

import (
	"encoding/json"
	"log"
	"net/http"

	// sms "github.com/baboooin/go_sms_ru"
	"github.com/pkg/errors"
)

func ValidateRegData(l RegisterRec) error {
	if l.Email == "" && l.Mobile == "" { //|| l.Password == ""
		return errors.New("Required fields empty")
	}

	if l.Email != "" {
		if !patternEmail.MatchString(l.Email) {
			return errors.New("Email not valid")
		}
	}

	if l.Mobile != "" {
		if !patternMobile.MatchString(l.Mobile) {
			return errors.New("Mobile not valid")
		}
	}

	if l.FistsName != "" {
		if !patternText.MatchString(l.FistsName) {
			return errors.New("First name not valid")
		}
	}

	if l.FistsName != "" {
		if !patternText.MatchString(l.FistsName) {
			return errors.New("First name not valid")
		}

	}

	if l.TaxNum != "" {
		if !patternTax.MatchString(l.TaxNum) {
			return errors.New("TaxID not valid")
		}

	}

	return nil
}

func register(w http.ResponseWriter, r *http.Request) {

	var l RegisterRec
	switch r.Method {
	case "GET":
		r.ParseForm()
		l.Email = r.FormValue("email")
		l.Mobile = r.FormValue("mobile")
		l.FistsName = r.FormValue("firstname")
		l.LastName = r.FormValue("lastname")
		l.TaxNum = r.FormValue("taxnum")
		l.Password = r.FormValue("password")
		// l.Password = r.FormValue("password")
	case "POST":
		if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
			Error(errors.New("Error query"), w)
			return
		}
	default:
		Error(errors.New("ErrorQuery"), w)
		return
	}

	if err := ValidateRegData(l); err != nil {
		Error(err, w)
		return
	}

	// if l.Mobile != "" {
	// 	twilio := NewTwillo(TwilioSID, TwilioToken)
	// 	verify, err := twilio.DoLookup(l.Mobile)
	// 	if err != nil {
	// 		Error(err, w)
	// 		return
	// 	}
	// 	log.Println(verify)
	// 	if !verify["valid"].(bool) {
	// 		Error(errors.New(verify["validation_errors"].(string)), w)
	// 		return
	// 	}

	// 	l.Mobile = verify["phone_number"].(string)
	// 	c := sms.SmsRU("26A73D0D-5116-8AF6-B031-2E9E6CD4483F")

	// 	// switch verify["country_code"].(string) {
	// 	// 	case "RU":
	// 	res, err := c.Call(l.Mobile)
	// 	if err != nil {
	// 		Error(err, w)
	// 		return
	// 	}

	// 	l.Mobile = res["code"].(string)
	// 	// 	default:

	// 	// }

	// }

	
	log.Println(l)

	// go SendPassword(l)

	json.NewEncoder(w).Encode(l)
	// return

	_, err := CreateUser(l)
	if err != nil {
		Error(err, w)
		return
	}

	// token, err := CreateToken(&u)
	// if err != nil {
	// 	Error(err, w)
	// 	return
	// }

	var reply SuccesReply
	reply.Code = 200
	reply.Message = "User successfuly created"
	// reply.Token = token

	json.NewEncoder(w).Encode(reply)

}
