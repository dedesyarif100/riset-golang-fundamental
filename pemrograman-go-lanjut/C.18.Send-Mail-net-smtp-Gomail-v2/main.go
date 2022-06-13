package main

import (
	"fmt"
	"log"
	// "strings"
	"net/smtp"
)

// const CONFIG_SMTP_HOST = "smtp.mailtrap.io"
// const CONFIG_SMTP_PORT = 2525
// const CONFIG_SENDER_NAME = "PT. Makmur Subur Jaya <emailanda@gmail.com>"
// const CONFIG_AUTH_EMAIL = "dedesyarif177@gmail.com"
// const CONFIG_AUTH_PASSWORD = "passwordemailanda"

// func sendMail(to []string, cc []string, subject, message string) error {
//     body := "From: " + CONFIG_SENDER_NAME + "\n" +
//         "To: " + strings.Join(to, ",") + "\n" +
//         "Cc: " + strings.Join(cc, ",") + "\n" +
//         "Subject: " + subject + "\n\n" +
//         message

//     auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
//     smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

//     err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
//     if err != nil {
//         return err
//     }

//     return nil
// }

func main() {
    from := "john.doe@example.com"
    user := "bbd715f1185866"
    password := "a38b50ef1efbd7"
    to := []string{
        "dedesyarif177@gmail.com",
    }

    addr := "smtp.mailtrap.io:2525"
    host := "smtp.mailtrap.io"

    msg := []byte("From: john.doe@example.com\r\n" +
        "To: dedesyarif177@gmail.com\r\n" +
        "Subject: Test mail\r\n\r\n" +
        "HALLO DEDE\r\n")

    auth := smtp.PlainAuth("", user, password, host)

    err := smtp.SendMail(addr, auth, from, to, msg)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Email sent successfully")
}