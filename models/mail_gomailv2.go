package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	"gopkg.in/gomail.v2"
)

// Both user list string or models.User list can be the recipient.
type MailDialer struct {
	Subject string
	Body    string
	To      User
}

var adminEmail, adminMailboxPassword,
	mailboxHost string
var mailboxPort int

type TemplateData map[string]string

func NewMailDialer(subject, body string, user User) *MailDialer {
	return &MailDialer{
		Subject: subject,
		Body:    body,
		To:      user,
	}
}

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

func (req *MailDialer) SendMail_gomailV2() error {

	mail := gomail.NewMessage()
	mail.SetHeader("From", adminEmail)

	// send by user list string
	mail.SetHeader("To", req.To.Email)

	mail.SetHeader("Subject", req.Subject)
	mail.SetBody("text/html", req.Body)

	dialer := gomail.NewDialer(mailboxHost, mailboxPort,
		adminEmail, adminMailboxPassword)

	if err := dialer.DialAndSend(mail); err != nil {
		log.Panic(err)
		return err
	}
	return nil
}
