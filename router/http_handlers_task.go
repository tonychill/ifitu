package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tonychill/ifitu/lib/utils"
)

func (r *routerImpl) handleCreateTask(c *fiber.Ctx) error {
	type something struct {
	}

	some := &something{}
	if err := utils.DecodeFiberRequest(c, some); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"request_id": reqId,
			"message":    fmt.Sprintf("Error decoding client's request: %s", err.Error()),
		})
	}
	// Your code here
	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"rules": "test",
	// })
	// Create a new Gmail API client
	// client, err := createGmailClient()
	// if err != nil {
	// 	log.Fatalf("Failed to create Gmail client: %v", err)
	// }

	// // Send email
	// err = sendEmail(client, "recipient@example.com", "Subject", "Body")
	// if err != nil {
	// 	log.Fatalf("Failed to send email: %v", err)
	// }

	// r.finImpl.CaptureFunds(c.Context(), &finSvc.CaptureFundsRequest{})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rules": "test",
	})
}

// func createGmailClient() (*gmail.Service, error) {
// 	// Load credentials from a file (replace with your own credentials file)
// 	credentialsFile := "path/to/credentials.json"
// 	config, err := google.ConfigFromJSON(credentialsFile, gmail.MailGoogleComScope)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load credentials: %v", err)
// 	}

// 	// Create a new OAuth2 token source
// 	tokenSource := config.TokenSource(context.Background(), &oauth2.Token{})

// 	// Create a new Gmail client
// 	client, err := gmail.NewService(context.Background(), oauth2.ReuseTokenSource(nil, tokenSource))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create Gmail client: %v", err)
// 	}

// 	return client, nil
// }

// func sendEmail(client *gmail.Service, recipient, subject, body string) error {
// 	// Create the email message
// 	message := gmail.Message{
// 		Raw: base64.URLEncoding.EncodeToString([]byte(
// 			fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", recipient, subject, body),
// 		)),
// 	}

// 	// Send the email
// 	_, err := client.Users.Messages.Send("me", &message).Do()
// 	if err != nil {
// 		return fmt.Errorf("failed to send email: %v", err)
// 	}

// 	return nil
// }
