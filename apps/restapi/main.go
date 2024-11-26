package main

import (
	"fmt"
	"log"
	"github.com/myusername/my-monorepo/pkg/sendGridGo"
)

func main() {
	emailData := sendGridGo.EmailData{
		Subject: "Test Email from App2",
		Email:   "anotherrecipient@example.com",
		Body:    "This is a test email from App2.",
		Name:    "App2 Sender",
	}

	htmlContent := "<html><body><h1>Email Body from App2</h1></body></html>"

	_, err := sendGridGo.SendGrid(emailData, htmlContent)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	fmt.Println("Email sent successfully from App2!")
}