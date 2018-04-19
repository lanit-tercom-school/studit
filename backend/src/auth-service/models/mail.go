package models

import (
	"crypto/tls"
	"fmt"
	"log"
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
		SmtpHttpsPort:  mailconf.String("https_port"),
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

func SendingRegistrationToken(pass string, Login string) *error {

	//build a message
	mail := Mail{}

	mail.senderId = mailConfig.SenderEmail
	mail.toId = Login
	mail.subject = "StudIT registration"
	mail.body = "Copy this registration code:\n\n"
	mail.pass = pass

	messageBody := mail.BuildMessage()
	smtpServer := SmtpServer{host: mailConfig.SmtpUrl, port: mailConfig.SmtpHttpsPort}

	log.Println(mailConfig)
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
		beego.Debug("Register error:" + err.Error())
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		beego.Debug("Register error:" + err.Error())
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		beego.Debug("Register error:" + err.Error())
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		beego.Debug("Register error:" + err.Error())
	}

	if err = client.Rcpt(mail.toId); err != nil {
		beego.Debug("Register error:" + err.Error())
	}

	// Data
	w, err := client.Data()
	if err != nil {
		beego.Debug("Register error:" + err.Error())
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		beego.Debug("Register error:" + err.Error())
	}

	err = w.Close()
	if err != nil {
		beego.Debug("Register error:" + err.Error())
	}

	client.Quit()
	return nil
}
