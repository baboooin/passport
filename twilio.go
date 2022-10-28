package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	auth  string
	sid   string
	token string
	http  *http.Client
}

func NewTwillo(sid string, token string) (client *Client) {
	data := fmt.Sprintf("%s:%s", sid, token)
	return &Client{
		auth:  fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(data))),
		http:  &http.Client{},
		sid:   sid,
		token: token,
	}
}

func (client *Client) DoLookup(number string) (data map[string]interface{}, err error) {

	url := fmt.Sprintf("https://lookups.twilio.com/v2/PhoneNumbers/%s", number)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) SendSMS(number string) error {

	return nil
}
