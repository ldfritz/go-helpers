package gmailhelper

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/joho/godotenv"
	"github.com/ldfritz/go-helpers/googleservices"
	"golang.org/x/net/context"
	gmail "google.golang.org/api/gmail/v1"
)

func Send() error {
	ctx := context.Background()
	env, err := godotenv.Read()
	if err != nil {
		return fmt.Errorf("Unable to load env: %v", err)
	}

	messageFilename := env["messagefile"]
	secretFilename := env["secretfile"]
	tokenFilename := env["tokenfile"]

	permissions := gmail.GmailSendScope

	ctx, config, token, err := googleservices.Authenticate(ctx, secretFilename, tokenFilename, permissions)
	if err != nil {
		return fmt.Errorf("Unable to authenticate session: %v", err)
	}

	svc, err := googleservices.Gmail(ctx, config, token)
	if err != nil {
		return fmt.Errorf("Unable to connect to service: %v", err)
	}

	message, nil := loadMessage(messageFilename)
	if err != nil {
		return fmt.Errorf("Unable to load message: %v", err)
	}

	_, err = svc.Users.Messages.Send("me", message).Do()
	if err != nil {
		return fmt.Errorf("Unable to send email: %v", err)
	}

	return nil
}

func loadMessage(filename string) (*gmail.Message, error) {
	rawMsg, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Cannot read %s: %v\n", filename, err)
	}

	var msg gmail.Message
	msg.Raw = base64.URLEncoding.EncodeToString(rawMsg)
	return &msg, nil
}
