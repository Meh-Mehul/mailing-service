package utils

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)
func SendMail(email, courseid string) error {
	from := os.Getenv("MAIL")
	password := os.Getenv("APP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := 587
	if email == "" {
		return fmt.Errorf("student_email is empty")
	}
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "[Update] Course pre-Registration Rejected")
	m.SetBody("text/plain", fmt.Sprintf(
		"Kindly be notified that your pre-registration for course %s has been rejected. Please contact admin in case of any issues.",
		courseid,
	))
	d := gomail.NewDialer(host, port, from, password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Email send failed:", err)
		return err
	}
	return nil
}
