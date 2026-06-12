package mailer

import (
	"bytes"
	"embed"
	"time"

	ht "html/template"
	tt "text/template"

	"github.com/wneessen/go-mail"
)

//go:embed "templates"
var templateFS embed.FS

type Mailer struct {
	client *mail.Client
	sender string
}

func New(host string, port int, username, password, sender string) (*Mailer, error) {

	client, err := mail.NewClient(
		host,
		mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithPort(port),
		mail.WithUsername(username),
		mail.WithPassword(password),
		mail.WithTimeout(5*time.Second),
	)

	if err != nil {
		return nil, err
	}

	return &Mailer{client: client, sender: sender}, nil

}

func (m *Mailer) Send(recipient, templateFile string, data any) error {

	textTemplate, err := tt.New("").ParseFS(templateFS, "templates/"+templateFile)

	if err != nil {
		return err
	}

	var subject = &bytes.Buffer{}
	err = textTemplate.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	var plainBody = &bytes.Buffer{}
	err = textTemplate.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return err
	}

	htmlTemplate, err := ht.New("").ParseFS(templateFS, "templates/"+templateFile)

	if err != nil {
		return err
	}

	var htmlBody = &bytes.Buffer{}

	err = htmlTemplate.ExecuteTemplate(htmlBody, "htmlBody", data)

	if err != nil {
		return err
	}

	msg := mail.NewMsg()

	err = msg.To(recipient)

	if err != nil {
		return err
	}

	err = msg.From(m.sender)

	if err != nil {
		return err
	}

	msg.Subject(subject.String())

	msg.SetBodyString(mail.TypeTextPlain, plainBody.String())

	msg.AddAlternativeString(mail.TypeTextHTML, htmlBody.String())

	return m.client.DialAndSend(msg)

}
