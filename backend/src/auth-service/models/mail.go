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
	senderId string
	toId     string
	subject  string
	body     string
	pass     string
}

type SmtpServer struct {
	host string
	port string
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

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
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

func SendingRegistrationToken(pass string, Login string) error {

	mail := Mail{}

	//build a message
	mail.senderId = mailConfig.SenderEmail
	mail.toId = Login
	mail.subject = "StudIT registration"
	mail.body = "Copy this registration code:\n\n"
	mail.pass = pass

	messageBody := mail.BuildMessage()
	smtpServer := SmtpServer{host: mailConfig.SmtpUrl, port: mailConfig.SmtpHttpsPort}

	//build an auth
	auth := smtp.PlainAuth("", mail.senderId, mailConfig.SenderPassword, smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		return err
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		return err
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		return err
	}

	if err = client.Rcpt(mail.toId); err != nil {
		return err
	}

	// Data
	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	client.Quit()
	return nil
}
