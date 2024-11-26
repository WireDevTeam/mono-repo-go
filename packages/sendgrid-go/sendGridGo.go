package sendGridGo

import (
	"fmt"
	"log"
	"os"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


type EmailData struct {
	Subject string `json:"subject"`
	Email   string `json:"email"`
	Body string `json:"body"`
	Name string `json:"name"`
}


// SendMail processes the incoming message and sends an email using SMTP
func SendGrid(emailData EmailData, htmlContent string) (EmailData, error) {
	// Process the message
	fmt.Printf("Received message: %s\n", emailData.Subject)
	
	validate(emailData)

	from := mail.NewEmail(emailData.Name, os.Getenv(("SENDGRID_SENDER")))
	subject := emailData.Subject
	to := mail.NewEmail("", emailData.Email)

	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
		return emailData, fmt.Errorf("failed to send email: %w", err) // Return a wrapped error
	} 

	switch response.StatusCode {
	case 202:
		fmt.Println("Email sent successfully:", response.StatusCode)
		return emailData, nil
	default:
		errMsg := fmt.Sprintf("unable to send mail: status code %d", response.StatusCode)
		return emailData, fmt.Errorf(errMsg) // Return an error with the status code
	}
}


func validate(emailData EmailData) error {
	// Validate environment variables
	if os.Getenv("SENDGRID_SENDER") == "" {
		return fmt.Errorf("missing SENDGRID_SENDER environment variable")
	}
	if os.Getenv("SENDGRID_API_KEY") == "" {
		return fmt.Errorf("missing SENDGRID_API_KEY environment variable")
	}

	// Additional email data validation can be added here
	if emailData.Email == "" || emailData.Subject == "" {
		return fmt.Errorf("invalid email data: email and subject are required")
	}

	return nil
}