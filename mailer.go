package main

import (
    "fmt"
    "strings"

    "crypto/tls"

    "gopkg.in/gomail.v2"
   // "github.com/aws/aws-sdk-go/service/ses"
    "github.com/badoux/checkmail"
)

func Send_Mail(ebook EBook, cfg Config) {
	fmt.Println("Sending email")
	if !ValidateAddress(ebook.Email) {
		return
	}

	d := gomail.NewDialer(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.User, cfg.SMTP.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.SMTP.Sender)
	m.SetHeader("To", ebook.Email)
	m.SetHeader("Subject", "A book for you")
	m.SetBody("text/html", "Happy birthday. Here goes a small <b>gift</b>!")
	m.Attach(ebook.Mobi)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}

}


func ValidateAddress(address string) bool  {

		err := checkmail.ValidateFormat(address)
		if err != nil {
			fmt.Println("format not valid")
			return false 
		}

	        components := strings.Split(address, "@")
    		_, domain := components[0], components[1]
		if domain !=  "kindle.com" {
			fmt.Println("domain is not kindle.com")
			return false
		}
	
		return true


}
