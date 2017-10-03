package googleservices

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	calendar "google.golang.org/api/calendar/v3"
	gmail "google.golang.org/api/gmail/v1"
	sheets "google.golang.org/api/sheets/v4"
	tasks "google.golang.org/api/tasks/v1"
)

func Calendar(ctx context.Context, config *oauth2.Config, token *oauth2.Token) (*calendar.Service, error) {
	svc, err := calendar.New(config.Client(ctx, token))
	if err != nil {
		return nil, fmt.Errorf("Unable to create client. %v", err)
	}
	return svc, nil
}

func Gmail(ctx context.Context, config *oauth2.Config, token *oauth2.Token) (*gmail.Service, error) {
	svc, err := gmail.New(config.Client(ctx, token))
	if err != nil {
		return nil, fmt.Errorf("Unable to create client. %v", err)
	}
	return svc, nil
}

func Sheets(ctx context.Context, config *oauth2.Config, token *oauth2.Token) (*sheets.Service, error) {
	svc, err := sheets.New(config.Client(ctx, token))
	if err != nil {
		return nil, fmt.Errorf("Unable to create client. %v", err)
	}
	return svc, nil
}

func Tasks(ctx context.Context, config *oauth2.Config, token *oauth2.Token) (*tasks.Service, error) {
	svc, err := tasks.New(config.Client(ctx, token))
	if err != nil {
		return nil, fmt.Errorf("Unable to create client. %v", err)
	}
	return svc, nil
}
