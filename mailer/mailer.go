package mailer

import (
	"log"
	"strings"

	"github.com/curiousz-peel/web-learning-platform-backend/config"
	gomail "gopkg.in/gomail.v2"
)

var MAIL config.EmailConfig

type Mailer struct {
	AddressTo string
}

func InitMail() {
	config, err := config.InitEmailConfig()
	if err != nil {
		log.Fatal("could not initialize config")
	}
	MAIL = *config
}

func (m *Mailer) SendRegistrationEmail() (err error) {
	registerMail := gomail.NewMessage()
	registerMail.SetHeader("From", MAIL.Address)
	registerMail.SetHeader("To", m.AddressTo)
	registerMail.SetHeader("Subject", "Registration Email")
	registerMail.SetBody("text/plain", "Please confirm your registration by clicking on the link below")

	var dispenser *gomail.Dialer
	splitToDomain := strings.Split(MAIL.Address, "@")
	mailFromType := strings.Split(splitToDomain[len(splitToDomain)-1], ".")[0]

	switch mailFromType {
	case "yahoo":
		dispenser = gomail.NewDialer("smtp.mail.yahoo.com", 587, MAIL.Address, MAIL.Password)
	default:
		log.Printf("unsuported FROM mail type: %v", mailFromType)
	}
	// dispenser.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := dispenser.DialAndSend(registerMail); err != nil {
		log.Printf("error while sending the registration mail to %v: %v", m.AddressTo, err)
	}

	return
}
