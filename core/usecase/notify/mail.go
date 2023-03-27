package notify

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	log "github.com/sirupsen/logrus"
)

//go:embed mail_template.html
var mailTemplateContent string
var mailTemplate *template.Template

type MailData struct {
	LookoutName string
}

func init() {
	tmpl, err := template.New("mail_template").Parse(mailTemplateContent)
	if err != nil {
		log.WithError(err).Fatal("Could not load mail template")
	}

	mailTemplate = tmpl
}

func sendMailNotification(value Notification) {
	from := "my_email@gmail.com"
	to := []string{"recipient@email.com"}
	smtpHost := "localhost"
	smtpPort := "1025"

	smtpAddress := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	fromHeader := fmt.Sprintf("From: %s\n", from)
	toHeader := fmt.Sprintf("To: %s\n", strings.Join(to, ","))
	subject := "Subject: Test email from Go!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := getMailBody(MailData{LookoutName: "test it"})
	msg := []byte(toHeader + fromHeader + subject + mime + body)

	// Send actual message
	err := smtp.SendMail(smtpAddress, nil, from, to, msg)
	if err != nil {
		log.WithError(err).Error("NOK")
		return
	}
}

func getMailBody(data MailData) string {
	bodyBuf := new(bytes.Buffer)
	if err := mailTemplate.Execute(bodyBuf, data); err != nil {
		log.WithError(err).Error("Could not load mail template")
		return ""
	}

	return bodyBuf.String()
}
