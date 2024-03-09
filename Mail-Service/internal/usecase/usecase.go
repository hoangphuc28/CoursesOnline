package usecase

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/Mail-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Mail-Service/internal/model"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"path/filepath"
)

type mailUsecase struct {
	cf *config.Config
}

func NewMailUsecase(cf *config.Config) *mailUsecase {
	return &mailUsecase{cf}
}

func (uc mailUsecase) SendEmail(email *model.Email, content *model.SendTokenContent) error {
	from := uc.cf.Email.AppEmail
	password := uc.cf.Email.AppPassword

	var body bytes.Buffer
	template, err := ParseTemplateDir("templates")
	if err != nil {
		return err
	}
	fmt.Println(content.Url)
	if err := template.ExecuteTemplate(&body, "verify.html", content); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email.DestMail)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil

}
func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
