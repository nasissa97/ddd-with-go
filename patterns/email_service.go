package patterns

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Interface to implement Infrasturce Service(MailChimp)
type EmailSender interface {
	SendEmail(ctx context.Context, to, title, body string) error
}

const emailURL = "https://someapp.com/api/1.0/messages/send"

// Infrasture that we want to use
type MailChimp struct {
	apiKey     string
	from       string
	httpClient http.Client
}

// Struct that make email request
type Recipient struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

type EmaiMessage struct {
	FromEmail string      `json:"from_email`
	Subject   string      `json:"subject"`
	Text      string      `json:"text"`
	To        []Recipient `json:"to"`
}

type MailChimpReqBody struct {
	Key     string      `json:"key"`
	Message EmaiMessage `json:"message"`
}

// Creates new instance of MailChimp Infrastructure Service
func NewMailChimp(apiKey, from string, httpClient http.Client) *MailChimp {
	return &MailChimp{apiKey: apiKey, from: from, httpClient: httpClient}
}

// Implemention of EmailSender for MailChimp
func (m MailChimp) SendEmail(ctx context.Context, to, title, body string) error {
	message := MailChimpReqBody{
		Key: m.apiKey,
		Message: EmaiMessage{
			FromEmail: m.from,
			Subject:   title,
			Text:      body,
			To: []Recipient{
				{
					to,
					"to",
				},
			},
		},
	}
	b, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshall body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, emailURL, bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	if _, err := m.httpClient.Do(req); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
