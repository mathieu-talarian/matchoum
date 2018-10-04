package mailing

import (
	"net/smtp"
	"fmt"
	"strings"
	"matcha/models"
)

var auth smtp.Auth
var host = "smtp.mailtrap.io"
var from = `"Matcha" <mmoullec@matcha.com>`;

type Mailtrap struct {
	Host     string
	Port     []string
	Username string
	Password string
	Auth     []string
}

var mailtrap = Mailtrap{
	"smtp.mailtrap.io",
	[]string{"25", "465", "2525"},
	"adfb57ba8a8ecc",
	"eb0fd49a79c987",
	[]string{"PLAIN", "LOGIN", "CRAM-MD5"},
}

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

func (s *Mailtrap) ServerName() string {
	return s.Host + ":" + s.Port[2]
}

func (mail *Mail) BuildMessage() []byte {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return []byte(message)
}

func (m Mail) Send() (error) {
	auth := smtp.PlainAuth("", mailtrap.Username, mailtrap.Password, mailtrap.Host)
	to := []string{"a@a.com"}

	return smtp.SendMail(mailtrap.ServerName(), auth, from, to, m.BuildMessage())
}

func NewMail(to []string, subject string, body string) (Mail) {
	mail := Mail{}
	mail.senderId = from
	mail.toIds = to
	mail.subject = subject
	mail.body = body
	return mail
}

func SendConfirmationEmail(u *models.User) (error) {
	m := NewMail([]string{u.Email},
		`Bienvenue sur Matcha`,
		fmt.Sprintf("Bienvenue sur Matcha, veuillez confirmer votre email\n\n%s\n\n",
			u.GenerateConfirmationUrl()))
	return m.Send()
}

func SendResetPasswordRequest(u *models.User) (error) {
	m := NewMail([]string{u.Email},
		`Reinitialisation de votre mot de passe`,
		fmt.Sprintf("Pour regenerer votre mot de passe cliquez sur ce lien\n\n%s\n", u.GenerateResetPasswordLink()))
	return m.Send()
}
