package utils

import (
	"crypto/rand"
	"crypto/tls"
	"github.com/hoangphuc28/CoursesOnline/User-Service/config"
	"gopkg.in/gomail.v2"
)

func SendToken(cf *config.Config, destMail string, data string) error {
	from := cf.Email.AppEmail
	//mật khẩu ứng dụng
	password := cf.Email.AppPassword
	////////////////////////////////

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", destMail)
	m.SetHeader("Subject", "Abc")
	m.SetBody("text/html", "http://127.0.0.1:5500/index.html?token="+data)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	////////////////////////////////////////////////////////////////
	//otp, err := GenerateOTP(6)
	//if err != nil {
	//	return "", err
	//}
	//
	//toEmailAddress := destMail
	//to := []string{toEmailAddress}
	//
	//host := "smtp.gmail.com"
	//port := "587"
	//address := host + ":" + port
	//
	//body :=
	//message := []byte(body)
	//
	//auth := smtp.PlainAuth("", from, password, host)
	//
	//if err := smtp.SendMail(address, auth, from, to, message); err != nil {
	//	return "", err
	//}

	return nil

}
func GenerateOTP(length int) (otp string, err error) {
	b := make([]byte, length)
	if _, err = rand.Read(b); err != nil {
		return "", err
	}
	for i := 0; i < length; i++ {
		otp += string(otpChars[int(b[i])%10])
	}
	return
}

const otpChars = "1234567890"
