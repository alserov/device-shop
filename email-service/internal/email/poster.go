package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	smtpAuthAddr   = "smtp.gmail.com"
	smtpServerAddr = "smtp.gmail.com:587"
)

type post struct {
	email    string
	password string
	name     string
}

type Poster interface {
	SendAuth(to string) error
	SendOrder(to string) error
}

func NewEmailManager(password, email, name string) Poster {
	return &post{
		email:    email,
		password: password,
		name:     name,
	}
}

func (p *post) send(title string, template []byte, to string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", p.name, p.email)
	e.Subject = title
	e.HTML = template
	e.To = []string{to}
	smtpAuth := smtp.PlainAuth("", p.email, p.password, smtpAuthAddr)

	return e.Send(smtpServerAddr, smtpAuth)
}

func (p *post) SendAuth(toEmail string) error {
	template := fmt.Sprintf(`
		<div style="background-color:white;height:500px;color:white">
			<h1 style="text-align:center">Hello, %s. You have successfully authorized.</h1>
			<p style="text-align:center">Click on this link to retrun to the page</p>
			<a href="http://127.0.0.1:/" style:"margin:5px">
				<button style="width:200px;height:60px;border-radius:10px;background-color:green;color:white;font-weight:bold;border:none">Back</button>
			</a>
		</div>
	`, toEmail)
	if err := p.send("Authorization", []byte(template), toEmail); err != nil {
		return err
	}
	return nil
}

func (p *post) SendOrder(toEmail string) error {
	template := fmt.Sprintf(`
		<p>New orde</p>
	`)
	if err := p.send("New order", []byte(template), toEmail); err != nil {
		return err
	}
	return nil
}
