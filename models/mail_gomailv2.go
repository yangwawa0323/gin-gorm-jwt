package models

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	"golang.org/x/term"
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

func Init() {
	// utils.LoadDotEnv()
	config := utils.InitConfig()
	adminEmail = config.Mailbox.AdminEmail

	adminMailboxPassword = config.Mailbox.AdminPassword
	if strings.Compare(adminMailboxPassword, "") == 0 {
		fmt.Printf("\nEnter admin mail box password (no echo):\n\n")

		// reader := bufio.NewReader(os.Stdin)
		// password, _, err := reader.ReadLine()

		bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			log.Panic(err)
			return
		}
		adminMailboxPassword = string(bytePassword)
	}
	mailboxHost = config.Mailbox.Host
	mailboxPort = int(config.Mailbox.Port)
	// mailboxPort, _ = strconv.Atoi(mailPort)
}

func (req *MailDialer) SendMail_gomailV2() error {
	Init()
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
