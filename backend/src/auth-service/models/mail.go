package models

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

type MailConfig struct {
	SmtpHttpsPort  string
	SmtpUrl        string
	SenderEmail    string
	SenderPassword string
}

var mailConfig MailConfig

type Mail struct {
	senderId       string
	senderPassword string
	toId           string
	subject        string
	body           string
	pass           string
}

func init() {
	mailconf, err := config.NewConfig("ini", "conf/mail.conf")
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
	mailConfig = MailConfig{
		SmtpHttpsPort:  mailconf.String("smtp_https_port"),
		SmtpUrl:        mailconf.String("smtp_host"),
		SenderEmail:    mailconf.String("sender_email"),
		SenderPassword: mailconf.String("sender_password"),
	}

}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toId) > 0 {
		message += fmt.Sprintf("To: %s\r\n", mail.toId)
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body + mail.pass

	return message
}

func (mail *Mail) Send(smtpUrl string, smtpPort string) error {
	//build an auth
	auth := smtp.PlainAuth("", mail.senderId, mail.senderPassword, smtpUrl)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpUrl,
	}

	var sendingEmailError = "403 Forbidden: "
	conn, err := tls.Dial("tcp", smtpUrl+":"+smtpPort, tlsconfig)
	if err != nil {
		beego.Error(sendingEmailError + err.Error())
		return err
	}

	client, err := smtp.NewClient(conn, smtpUrl)
	if err != nil {
		beego.Error(sendingEmailError + err.Error())
		return err
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		beego.Error(sendingEmailError + err.Error())
		return err
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		beego.Error(sendingEmailError + err.Error())
		return err
	}

	if err = client.Rcpt(mail.toId); err != nil {
		beego.Error(sendingEmailError + err.Error())
		return err
	}

	// Data
	w, err := client.Data()
	if err != nil {
		beego.Error(sendingEmailError + err.Error())
		return err
	}

	_, err = w.Write([]byte(mail.BuildMessage()))
	if err != nil {
		beego.Error(sendingEmailError + err.Error())
		return err
	}

	err = w.Close()
	if err != nil {
		beego.Error(sendingEmailError + err.Error())
		return err
	}

	client.Quit()
	return nil

}

func SendingRegistrationToken(pass string, Login string) error {

	mail := Mail{}

	//build a message
	mail.senderId = mailConfig.SenderEmail
	mail.toId = Login
	mail.subject = "StudIT registration"
	mail.body = "Copy this registration code:\n\n"
	mail.pass = pass
	mail.senderPassword = mailConfig.SenderPassword

	err := mail.Send(mailConfig.SmtpUrl, mailConfig.SmtpHttpsPort)
	return err
}
