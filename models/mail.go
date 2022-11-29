package models

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	"gopkg.in/gomail.v2"
)

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

var adminEmail, adminMailboxPassword,
	mailboxHost,
	activateUrl, secret string
var mailboxPort int
var auth smtp.Auth

const mailTemplateFile = "./templates/activate-mail.html"

type TemplateData map[string]string

func init() {
	utils.LoadDotEnv()
	adminEmail = os.Getenv("ADMIN_EMAIL")

	adminMailboxPassword = os.Getenv("ADMIN_MAILBOX_PASSWORD")
	if strings.Compare(adminMailboxPassword, "") == 0 {
		fmt.Printf("\nEnter admin mail box password\n\n")
		reader := bufio.NewReader(os.Stdin)
		password, _, err := reader.ReadLine()
		if err != nil {
			log.Panic(err)
			return
		}
		adminMailboxPassword = string(password)
	}
	mailboxHost = os.Getenv("MAILBOX_HOST")
	mailPort := os.Getenv("MAILBOX_PORT")
	mailboxPort, _ = strconv.Atoi(mailPort)
}

func SendMail_gomailV2() {

	var buf *bytes.Buffer = new(bytes.Buffer)
	templateData := GenerateTemplateData()
	tmpl, err := template.ParseGlob(mailTemplateFile)

	if err != nil {
		log.Panic("cannot parse the template file")
		return
	}
	if err := tmpl.Execute(buf, templateData); err != nil {
		log.Panic("can not execute applies a parsed template to specified data object")
		return
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", adminEmail)
	mail.SetHeader("To", "yangwawa0323@163.com")
	mail.SetHeader("Cc", "qin49@126.com")
	mail.SetHeader("Subject", "Welcome to register 51cloudclass.com")

	mail.SetBody("text/html", buf.String())

	dialer := gomail.NewDialer(mailboxHost, mailboxPort,
		adminEmail, adminMailboxPassword)

	if err := dialer.DialAndSend(mail); err != nil {
		log.Panic(err)
	}
}

func Send() {

	auth = smtp.PlainAuth("", adminEmail, adminMailboxPassword, mailboxHost)

	templateData := GenerateTemplateData()

	req := NewRequest([]string{"yangwawa0323@163.com", "12238747@qq.com"},
		"Welcome to register 51cloudclass.com",
		"",
	)
	if err := req.ParseTemplate(mailTemplateFile, templateData); err != nil {
		log.Fatal("can not read the template, ", err.Error())
	}

	ok, _ := req.SendEmail()
	utils.Debug(strconv.FormatBool(ok))

}

func GenerateTemplateData() *TemplateData {
	var templateData TemplateData = make(TemplateData)
	templateData["ActivateString"] = fmt.Sprintf("%s%s", activateUrl, secret)
	return &templateData
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		from:    adminEmail,
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (req *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\n"
	mime += "Content-Type: text/plain; charset=\"UTF-8\";\n\n"

	subject := "Subject: " + req.subject + "!\n"
	msg := []byte(subject + mime + "\n" + req.body)

	if err := smtp.SendMail(
		fmt.Sprintf("%s:%d", mailboxHost, mailboxPort),
		auth,
		adminEmail,
		req.to,
		msg); err != nil {
		return false, err
	}
	return true, nil
}

func (req *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}

	req.body = buf.String()
	return nil
}
