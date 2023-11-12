package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"os"
)

const (
	smtpAuthAddr   = "smtp.gmail.com"
	smtpServerAddr = "smtp.gmail.com:587"
)

func SendEmail(to string) error {
	var (
		sender = os.Getenv("SENDER_NAME")
		mail   = os.Getenv("SENDER_EMAIL")
		pass   = os.Getenv("EMAIL_PASSWORD")
	)

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender, mail)
	e.Subject = "Authorization"
	e.HTML = []byte(fmt.Sprintf(`
		<div style="background-color:#3b2e3e;height:500px;align-items:center;color:white;margin:30px,0">
			<h1 style="text-align:center">Hello, %s. You have successfully authorized.</h1>
			<p style="text-align:center">Click on this link to retrun to the page</p>
			<a href="http://127.0.0.1:/" style:"margin:5px">
				<button style="width:200px;height:60px;border-radius:10px;background-color:#601473;color:white;font-weight:bold;border:none">Back</button>
			</a>
		</div>
	`, to))
	e.To = []string{to}
	smtpAuth := smtp.PlainAuth("", mail, pass, smtpAuthAddr)

	return e.Send(smtpServerAddr, smtpAuth)
}
