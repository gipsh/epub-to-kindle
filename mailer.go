package main

import (
    "fmt"

    "crypto/tls"

    "gopkg.in/gomail.v2"
   // "github.com/aws/aws-sdk-go/service/ses"
)

func Send_Mail(ebook EBook, cfg Config) {
	fmt.Println("Sending email")

	d := gomail.NewDialer(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.User, cfg.SMTP.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "robot@ubiguard.com")
	m.SetHeader("To", ebook.Email)
	m.SetHeader("Subject", "A book for you")
	m.SetBody("text/html", "Happy birthday. This is your party. This a small <b>gift</b>!")
	m.Attach(ebook.Mobi)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

