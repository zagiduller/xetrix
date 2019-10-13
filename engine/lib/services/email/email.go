package email

import (
	"github.com/olebedev/config"
	"log"
	"net/smtp"
)

type Sender struct {
	hostport string
	auth     smtp.Auth
	from     string
}

func NewSender(cfg *config.Config) *Sender {
	s := new(Sender)
	// Set up authentication information.
	s.auth = smtp.PlainAuth(
		"",
		cfg.UString("username"),
		cfg.UString("password"),
		cfg.UString("host"),
	)
	s.hostport = cfg.UString("host") + cfg.UString("port")
	s.from = cfg.UString("from")
	return s
}

func (s *Sender) Send(addr string) error {
	log.Printf("Email: Send email from %s to %s", s.from, addr)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	return smtp.SendMail(
		s.hostport,
		s.auth,
		s.from,
		[]string{addr},
		[]byte("This is the email body."),
	)
}
