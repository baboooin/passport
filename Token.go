package main

import (
	"encoding/json"
	"net/http"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
)

func CreateToken(user *UserRecord) (string, error) {

	jUser, err := json.Marshal(&user)
	if err != nil {
		return "", err
	}
	pubkey, err := LoadPublicKey()
	if err != nil {
		return "", err
	}
	encode, err := jwe.Encrypt(jUser, jwe.WithKey(jwa.RSA1_5, pubkey))
	if err != nil {
		return "", err
	}

	return string(encode), err
}

func VerifyToken(Token string) ([]byte, error) {
	// decode
	kp, err := LoadPrivateKey()
	if err != nil {
		return nil, err
	}
	decode, err := jwe.Decrypt([]byte(Token), jwe.WithKey(jwa.RSA1_5, kp))
	return decode, err
}


func getToken (w http.ResponseWriter, r *http.Request) {
	var sr *SuccesReply 
token := "token"
sr.Code = 200
sr.Token = token
w.WriteHeader(http.StatusOK)	
json.NewEncoder(w).Encode(sr)
}


