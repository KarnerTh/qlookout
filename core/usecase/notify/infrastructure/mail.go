package infrastructure

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/usecase/notify"
)

//go:embed mail_template.html
var mailTemplateContent string

type mailNotifier struct {
	fromAddress  string
	smtpHost     string
	smtpPort     string
	mailTemplate *template.Template
}

func NewMailNotifier(fromAddress string, smtpHost string, smtpPort string) notify.Notifier {
	tmpl, err := template.New("mail_template").Parse(mailTemplateContent)
	if err != nil {
		log.WithError(err).Fatal("Could not load mail template")
	}

	return mailNotifier{
		fromAddress:  fromAddress,
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		mailTemplate: tmpl,
	}
}

func (n mailNotifier) Send(value notify.Notification) error {
	to := []string{"recipient@email.com"}
	smtpAddress := fmt.Sprintf("%s:%s", n.smtpHost, n.smtpPort)

	fromHeader := fmt.Sprintf("From: %s\n", n.fromAddress)
	toHeader := fmt.Sprintf("To: %s\n", strings.Join(to, ","))
	subject := "Subject: Test email from Go!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := getMailBody(n.mailTemplate, value)
	msg := []byte(toHeader + fromHeader + subject + mime + body)

	err := smtp.SendMail(smtpAddress, nil, n.fromAddress, to, msg)
	if err != nil {
		log.WithError(err).Error("Could not send mail")
		return err
	}

	return nil
}

func getMailBody(template *template.Template, data notify.Notification) string {
	bodyBuf := new(bytes.Buffer)
	if err := template.Execute(bodyBuf, data); err != nil {
		log.WithError(err).Error("Could not load mail template")
		return ""
	}

	return bodyBuf.String()
}
