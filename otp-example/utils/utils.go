package utils

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

func SendEmail(toEmail string, code int) error {

	m := gomail.NewMessage()
	fromMail := "wpapp23@outlook.com"
	fromPass := "Appsecret1."

	// Set E-Mail sender
	m.SetHeader("From", fromMail)

	// Set E-Mail receivers
	m.SetHeader("To", toEmail)

	// Set E-Mail subject

	body := "Doğrulama kodunuz:" + strconv.Itoa(code)
	subject := "Hesabınızı doğrulamak için gelen doğrulama e-postası"

	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", body)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.office365.com", 587, fromMail, fromPass)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func RandNumber(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(end-start) + start
	fmt.Println("code:", code)
	return code
}

func GetInput() string {

	fmt.Print("Enter code: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	return input
}
