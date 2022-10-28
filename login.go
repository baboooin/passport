package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func login(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var l LoginRec
	r.ParseForm()
	// l.Login, l.Password, ok = r.BasicAuth()
	// if !ok {
	log.Println(r.Method)
	switch r.Method {
	case "GET":
		r.ParseForm()
		// if r.FormValue("login") == "" || r.FormValue("password") == "" {
		// 	Error(nil, "ErrorQuery", w)
		// 	return
		// }

		l.Login = r.FormValue("login")
		l.Password = r.FormValue("password")

	case "POST", "PUT":
		{
			bodyBytes, err := postBody(r)
			if err != nil {
				Error(err, w)
			}
			x, _ := json.Marshal(bodyBytes)
			err = json.Unmarshal(x, &l)
			if err != nil {
				Error(errors.New("Test"), w)
			}
		}
	default:
		Error(errors.New("ErrorQuery"), w)
		return
	}

	if patternPassword.MatchString(l.Password) && (patternEmail.MatchString(l.Login) || patternMobile.MatchString(l.Login)) {

		if user, err := CheckUser(l); err != nil {
			Error(err, w)
		} else {
			if token, err := CreateToken(user); err != nil {
				Error(err, w)
			} else {
				var sr SuccesReply
				w.WriteHeader(http.StatusOK)
				sr.Code = 200
				sr.Message = "Login success"
				sr.Token = token
				json.NewEncoder(w).Encode(sr)
			}
		}
	} else {
		Error(errors.New("Auth error"), w)
	}
}
