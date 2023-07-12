package services

import (
	"github.com/3dw1nM0535/nyatta/config"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/resendlabs/resend-go"
)

type MailingServices struct {
	queries *sqlStore.Queries
	client  *resend.Client
}

// Enforce Mailing service interface
var _ interfaces.Mailing = &MailingServices{}

func NewMailingService(queries *sqlStore.Queries, cfg config.EmailConfig) *MailingServices {
	client := resend.NewClient(cfg.Apikey)
	return &MailingServices{queries: queries, client: client}
}

// SaveMailing saves an email to the database
func (m *MailingServices) SaveMailing(email string) (*model.Status, error) {
	// Mail exists
	exists, err := m.queries.MailingExists(ctx, email)
	if err != nil {
		return nil, err
	}
	// Save mail
	if !exists {
		_, err := m.queries.SaveMail(ctx, email)
		if err != nil {
			return nil, err
		}
		return &model.Status{Success: "okay"}, nil
	} else {
		return &model.Status{Success: "okay"}, nil
	}
}

func (m *MailingServices) SendEmail(to []string, from, subject, body string) error {
	params := &resend.SendEmailRequest{
		To:      to,
		From:    from,
		Subject: subject,
		Text:    body,
	}
	_, err := m.client.Emails.Send(params)
	if err != nil {
		return err
	}
	return nil
}

func (m MailingServices) ServiceName() string {
	return "Mailing"
}
