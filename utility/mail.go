package utility

import (
	"crypto/tls"
	"fmt"
	EmailTemplate "gosql/template/email"
	"log"
	"net/smtp"
	"strings"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Student App <galihkur29@gmail.com>"
const CONFIG_AUTH_EMAIL = "galihkur29@gmail.com"
const CONFIG_AUTH_PASSWORD = "aqogepbxmthlljpl"

func Send(target []string, target_cc []string, activation string) {

	gom := gomail.NewDialer(CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)
	gom.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// gom.SSL = true

	subject := "Activation user"
	data := activation
	message := EmailTemplate.ActivateUsers(data)

	send := gomail.NewMessage()
	send.SetHeader("From", CONFIG_SENDER_NAME)
	send.SetHeader("To", target...)
	send.SetHeader("Subject", subject)
	send.SetBody("text/html", message)

	// send
	err := gom.DialAndSend(send)
	if err != nil {
		CreateLog(err)
		log.Println(err)

	}
	log.Println("Mail sent!")

}

// func Send(target []string, target_cc []string, activation string) {

// 	to := target
// 	cc := target_cc
// 	subject := "Activation user"
// 	data := activation
// 	message := EmailTemplate.ActivateUsers(data)

// 	err := sendMail(to, cc, subject, message)
// 	if err != nil {
// 		CreateLog(err)
// 	}

// 	log.Println("Mail sent!")

// }

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		CreateLog(err)
		return err
	}

	return nil
}
