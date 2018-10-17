package main

import (
	"crypto/tls"
	"log"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func main() {
	// send("hello there")
	gomail_sendmail()
}

func gomail_sendmail() {
	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

func send(body string) {
	from := "talovenit1987@gmail.com"
	pass := "0835751891"
	to := "takissnit1987@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
