package communication

import (
	"crypto/tls"
	"go_worlder_system/consts"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var (
	d = newDialer()
)

func newDialer() *gomail.Dialer {
	user := os.Getenv(consts.EMAIL_FROM)
	host := os.Getenv(consts.EMAIL_HOST)
	password := os.Getenv(consts.EMAIL_PASSWORD)
	d := gomail.NewDialer(
		host,
		587,
		user,
		password,
	)
	d.TLSConfig = &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: true,
	}
	return d
}

func SendMail(to, sub, message string) error {
	user := os.Getenv(consts.EMAIL_FROM)
	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", to)
	m.SetHeader("Subject", sub)
	m.SetBody("text/plain", message)
	err := d.DialAndSend(m)
	if err != nil {
		log.Error(err)
	}
	return err
}
