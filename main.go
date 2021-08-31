package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	gomail "gopkg.in/mail.v2"
)

//this is the current email
var current Email

func EmailRequest(w http.ResponseWriter, r *http.Request) {
  //This si suppsoed to get the info from the http request
  var e Email

  err := json.NewDecoder(r.Body).Decode(&e)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

}

func handleRequests(){
  http.HandleFunc("/email", EmailRequest)
  log.Fatal(http.ListenAndServe(":10000", nil))
}


func main() {
  //this method sends the email
  m := gomail.NewMessage()

  // Set E-Mail sender
  m.SetHeader("From", current.Sender)

  // Set E-Mail receivers
  m.SetHeader("To", current.Reciever)

  // Set E-Mail subject
  m.SetHeader("Subject", current.Subject)

  // Set E-Mail body. You can set plain text or html with text/html
  m.SetBody("text/plain", current.BodyText)

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

