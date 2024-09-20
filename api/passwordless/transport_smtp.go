package passwordless

import (
	"context"
	"fmt"
	"net/smtp"
	"os"
)

// SMTPConfig holds the configuration for SMTPTransport.
type SMTPConfig struct {
	UseSSL   bool
	Addr     string
	From     string
	Auth     smtp.Auth
}

// SMTPTransport delivers a user token via e-mail.
type SMTPTransport struct {
	config SMTPConfig
}

// NewSMTPTransport returns a new transport capable of sending emails via SMTP.
func NewSMTPTransport(config SMTPConfig) *SMTPTransport {
	return &SMTPTransport{config}
}

// SendToken sends a token using the log transport.
func (s *SMTPTransport) SendToken(ctx context.Context, email, token string, tokenType TokenType) error {
	baseURL := os.Getenv("APP_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	subject, body := composeEmail(token, tokenType, baseURL)
	to := []string{email}
	msg := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
	
	if s.config.UseSSL {
		return sendMailSSL(s.config.Addr, s.config.Auth, s.config.From, to, msg)
	} else {
		return smtp.SendMail(s.config.Addr, s.config.Auth, s.config.From, to, msg)
	}
}

func composeEmail(token string, tokenType TokenType, baseURL string) (subject, body string) {
	switch tokenType {
	case TokenTypeString:
		magicLink := fmt.Sprintf("%s/authenticate?token=%s", baseURL, token)
		subject = "Your Magic Link"
		body = fmt.Sprintf("Hello,\n\nPlease use the following link to login: %s", magicLink)
	case TokenTypePin:
		subject = "Your PIN"
		body = fmt.Sprintf("Hello,\n\nYour PIN is: %s", token)
	}
	return
}

func sendMailSSL(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	return nil
}
