package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	gomail "gopkg.in/mail.v2"
)

func setemail(sender string, reciever string, subject string, body string) {
	//this method sends the email
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", sender)

	// Set E-Mail receivers
	m.SetHeader("To", reciever)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", body)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "from@gmail.com", "qborwrttzdfnvyvu")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// return Email
}

func getSender(r *http.Request)string {
	vars := mux.Vars(r)
	sender := (vars["sender"])
	return sender
}

func getReciever(r *http.Request) string{
	vars := mux.Vars(r)
	rec := (vars["recieve"])
	return rec
}

func getSubject(r *http.Request) string{
	vars := mux.Vars(r)
	rec := (vars["subject"])
	return rec
}

func getBody(r *http.Request)string {
	vars := mux.Vars(r)
	rec := (vars["body"])
	return rec
}

func GetThatMail(rw http.ResponseWriter, r *http.Request) {
	send := getSender(r)
	to := getReciever(r)
	sub := getSubject(r)
	bod:= getBody(r)
	
	setemail(send, to, sub, bod)

	fmt.Println("Should Send Email IG")
}