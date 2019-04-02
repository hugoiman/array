package email

import (
  "fmt"
  // "gopkg.in/gomail.v2"
)

// const CONFIG_SMTP_HOST = "smtp.gmail.com"
// const CONFIG_SMTP_PORT = 587
// const CONFIG_EMAIL = "nanonymoux@gmail.com"
// const CONFIG_PASSWORD = "kudaponi10080"

func SendEmail(to, subject, message string) {
  fmt.Println(to)
  fmt.Println(subject)
  fmt.Println(message)
  // mailer := gomail.NewMessage()
  // mailer.SetHeader("From", CONFIG_EMAIL)
  // mailer.SetHeader("To", to)
  // mailer.SetHeader("Subject", subject)
  // mailer.SetBody("text/html", message)
  //
  // dialer := gomail.NewDialer(CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_EMAIL, CONFIG_PASSWORD)
  //
  // err := dialer.DialAndSend(mailer)
  // checkErr(err)
}
