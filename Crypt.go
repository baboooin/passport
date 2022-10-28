package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/lestrrat-go/jwx/v2/jwk"
)



func LoadPrivateKey() (*rsa.PrivateKey, error) {
	raw, err := ioutil.ReadFile(PrivateKeyFile)
	if err != nil {
		return nil, err
	}
	privPem, _ := pem.Decode(raw)

	parsedKey, err := x509.ParsePKCS8PrivateKey(privPem.Bytes)
	// var privateKey *rsa.PrivateKey
	privateKey, _ := parsedKey.(*rsa.PrivateKey)
	return privateKey, err
}



func LoadPublicKey() (jwk.Key, error) {
	privKey, err := LoadPrivateKey()
	if err!=nil {return nil,err}
	pubKey, err := jwk.FromRaw(privKey.PublicKey)
	return pubKey, err
}