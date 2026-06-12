package mailer

import (
	"embed"
	"time"

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

func (m *Mailer) ale(templateFile string) {

	textTemplate, err := tt.New("").ParseFS(templateFS, "templates/"+templateFile)

	if err != nil {

	}
}
