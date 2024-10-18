package main

import (
	"fmt"
	"os"

	spaceapi "github.com/Ozii-cr/notify-space/pkg/space_api"
	spacemessage "github.com/Ozii-cr/notify-space/pkg/space_message"
)

func Run() error {
	githubEvent := os.Getenv("GITHUB_EVENT_NAME")
	githubRepo := os.Getenv("GITHUB_REPOSITORY")
	spaceWebhookUrl := os.Getenv("INPUT_SPACE_WEBHOOK_URL")
	messageType := os.Getenv("INPUT_MESSAGE_TYPE")

	client := spaceapi.NewClient(spaceWebhookUrl)

	var message spacemessage.Message

	switch messageType {
	case "plain":
		message = spacemessage.NewPlainMessage(fmt.Sprintf("New activity in %s: %s event occurred", githubRepo, githubEvent))

	default:
		return fmt.Errorf("invalid message type specified")
	}

	err := client.SendMessage(message)
	if err != nil {
		return fmt.Errorf("error sending message: %v", err)
	}

	fmt.Println("Message sent successfully to Space")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Printf("Fatal error: %v\n", err)
		os.Exit(1)
	}
}
