package main

import (
	//  "github.com/emersion/go-smtp"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net"
	"net/smtp"
	"strings"
	"time"

	dkim "github.com/toorop/go-dkim"
)


func SendMailTempotaryPassword(rcpt string, nameRcpt string, subject string, templatePath string, password string) error {

	PrivateKey := `LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWEFJQkFBS0JnUURrSlF5aVV2bkU4K0RnSTloSC8veWV6RTRjSUhOR3dOd2J4SGF6bE13Vm82djA3QWtjWHUrVHloZU8KTllCamVpYjFzN1FwS3ZKK3hVWVJETk54SHMyTE5EQUE4cnhQdm5QMDJ4SWdHOUlyb2Jxa05BeVdweFRJU1RiNk5FczdBd2dNCjBsbmZhampsa2cxdnVjNE45UTU4WmxvVWNYWU1lczNoQzUyL1lKbGc0UUlEQVFBQkFvR0FVYmhmclNsUnBGWERJQ1pXeTN1TApUb3BhRkVNZlo2R2ZkaWJLYWdzTGkxakVPSWZ2djRUV0JqY21kdDc5TFNUTkhjZVA2Z3NqTDN1VVhIa2VwKzlqcWlNUW5oTDIKc1dGNmN5dFRNNHY0VnRQS1BUZTF6cnVKZ3QxZjlIVExSVkV6WFdUZlpNb0FZMFprajluNGZ6Mk1GbnRQWmlLUHVvYlQ2K3NGCjcrY2dOeEVDUVFEMWo0czNac3VBbUNnV2o2dHd2RlgwMWF6Y21XeTNKamk0K1lWcWVET2VhMlJWdG1QYkQwektZRGptYzhaawpTc0t1R0YxWjZxcWtxUUhqdldFR0FoeUZBa0VBN2RmM1I0aTVSWmVkQnVzWHUyNzNHVFJOTXUyUjc4RTJaMU5kMjRFeDVYdDkKS0dpMURCYzZ1WkVwekloNHhQUU9uTTgveWhCNnVoK21uNnY5MmxnZnJRSkFMN3RXUVUrZThRRWlrYmdkaUExajIvS0k1bHlBClVGMkNteG9OZi9PYmZRaTgrUmc1OHNQWGxtTDd6SGZtc1dvQ3k1cDYwdGFWa3Vvekx0U0ovb1A4alFKQkFKM2p5b0l3WDEyMQpna1ovZG1lMUNGQXhDamFPdzF4MzFSZk9uZllEMUEzVlpvYVd5K2xVMm1VcDJxZXV5UTFtUHZVV0YvQ1o4Lzd1MS93VGZ4ZWQKZkprQ1FGTTRudS9RTFBWUm1YeitYUWh2d0wrdS9sQlJ4WnBnVjJzc1psSFhJRTJoemNZZjUyQk1yQ0puUGZycjh5dTlQc3g3CmcxdHQ4M2J5NGxHTm84M2ZlTjg9Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==`

	from := "ice@micro.bz"
	fromName := "Passport"

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Date:%s\n", time.UnixDate)))
	body.Write([]byte(fmt.Sprintf("From: %s <%s>\n", fromName, from)))
	body.Write([]byte(fmt.Sprintf("To: %s <%s>\n", nameRcpt, rcpt)))
	body.Write([]byte(fmt.Sprintf("Subject:%s\n%s\n\n", subject, mimeHeaders)))

	t.Execute(&body, struct {
		Name     string
		Password string
	}{
		Name:     "l.FistsName l.LastName",
		Password: "l.Password",
	})

	options := dkim.NewSigOptions()
	options.PrivateKey = []byte(PrivateKey)
	options.Domain = "micro.bz"
	options.Selector = "mail"
	options.SignatureExpireIn = 3600
	options.BodyLength = 50
	options.Headers = []string{"from", "date", "mime-version", "received", "received"}
	options.AddSignatureTimestamp = true
	options.Canonicalization = "relaxed/relaxed"

	bytes := body.Bytes()
	err = dkim.Sign(&bytes, options)

	fmt.Println(err)

	fmt.Println(string(bytes))

	name := strings.Split(rcpt, "@")[1]
	mx, _ := net.LookupMX(name)
	// addr := strings.Split(mx[0].Host, ".")
	addr := mx[0].Host
	c, err := smtp.Dial(addr + ":25")

	if err != nil {
		err = errors.New("Dial " + err.Error())
		return err
	}

	err = c.Hello("micro.bz")
	if err != nil {
		err = errors.New("HELO " + err.Error())
		return err
	}

	// err =c.Verify(rcpt)
	// if err != nil {
	// 	err = errors.New("Verify: " + err.Error())
	// 	return err
	// }
	err = c.Mail("ice@micro.bz")
	if err != nil {
		err = errors.New("From " + err.Error())
		return err
	}

	err = c.Rcpt(string(rcpt))
	if err != nil {
		err = errors.New("TO " + err.Error())
		return err
	}

	w, err := c.Data()
	if err != nil {
		err = errors.New("DATA " + err.Error())
		return err
	}
	_, err = w.Write(bytes)
	if err != nil {
		err = errors.New("Write data " + err.Error())
		return err
	}
	_, err = w.Write([]byte("."))
	if err != nil {
		err = errors.New("write dot" + err.Error())
		return err
	}

	err = c.Quit()
	if err != nil {
		err = errors.New("Close " + err.Error())
		return err
	}

	err = c.Close()
	if err != nil {
		err = errors.New("Close " + err.Error())
		return err
	}

	return nil
}
