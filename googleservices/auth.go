package googleservices

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func Authenticate(ctx context.Context, secretFilename, tokenFilename, permissions string) (context.Context, *oauth2.Config, *oauth2.Token, error) {
	config, err := getConfig(secretFilename, permissions)
	if err != nil {
		return ctx, nil, nil, fmt.Errorf("Unable to load config. %v", err)
	}

	token, err := getToken(config, tokenFilename)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Unable to load token. %v", err)
	}

	return ctx, config, token, nil
}

func getConfig(secretFilename, permissions string) (*oauth2.Config, error) {
	secret, err := ioutil.ReadFile(secretFilename)
	if err != nil {
		return nil, fmt.Errorf("Unable to read secret file: %s. %v", secretFilename, err)
	}

	config, err := google.ConfigFromJSON(secret, permissions)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse secret. %v", err)
	}

	return config, nil
}

func getToken(config *oauth2.Config, tokenFilename string) (*oauth2.Token, error) {
	f, err := os.Open(tokenFilename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)

	if err != nil {
		log.Printf("Unable to load token file: %s. %v\n", tokenFilename, err)
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		fmt.Printf("Authentication URL:\n%v\nAuthentication code: ", authURL)

		var code string
		if _, err := fmt.Scan(&code); err != nil {
			return nil, fmt.Errorf("Unable to read authentication code. %v", err)
		}

		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			return nil, fmt.Errorf("Unable to download token. %v", err)
		}

		log.Printf("Saving token to file: %s\n", tokenFilename)
		f, err := os.OpenFile(tokenFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return nil, fmt.Errorf("Unable to save token file: %s. %v", tokenFilename, err)
		}
		defer f.Close()
		json.NewEncoder(f).Encode(token)
	}
	return token, nil
}
