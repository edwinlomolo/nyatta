package services

import (
	"github.com/3dw1nM0535/nyatta/config"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	"github.com/resendlabs/resend-go"
	"github.com/sirupsen/logrus"
)

const (
	newUserEmail     = "<p>Welcome to Nyatta! I'm excited to have you onboard!<br /><br /> How best can you describe your house hunting experience? Exactly<br /><br /> I'm working hard to make Nyatta the best place to find, move in, and share your favorite places with friends. You may encounter or face some issues but I'm able to track them as they happen. However, if you see anything of concern, I'm happy to have a chat.<br /><br /> Please reach out if you have any questions or feedback for me. <br /><br /> <strong>Regards,</strong><br /><br /> Edwin Lomolo.</p>"
	newPropertyEmail = "<p>I've received your listing! I'm always pushing product updates weekly and you will receive an email about what has changed/what's new. I'm always looking into how I can make your experience using Nyatta better. Don't hesitate to reach out!<br /><br /> <strong>Regards,</strong><br /><br /> Edwin Lomolo.</p>"
)

type SendEmail func(email []string, from, subject, body string) error

type MailingServices struct {
	queries *sqlStore.Queries
	client  *resend.Client
	log     *logrus.Logger
}

// Enforce Mailing service interface
var _ interfaces.Mailing = &MailingServices{}

func NewMailingService(queries *sqlStore.Queries, cfg config.EmailConfig, logger *logrus.Logger) *MailingServices {
	client := resend.NewClient(cfg.Apikey)
	return &MailingServices{queries: queries, client: client, log: logger}
}

// SaveMailing saves an email to the database
func (m *MailingServices) SaveMailing(email string) (*model.Status, error) {
	// Mail exists
	exists, err := m.queries.MailingExists(ctx, email)
	if err != nil {
		m.log.Errorf("%s:%v", m.ServiceName(), err)
		return nil, err
	}
	// Save mail
	if !exists {
		_, err := m.queries.SaveMail(ctx, email)
		if err != nil {
			m.log.Errorf("%s:%v", m.ServiceName(), err)
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
		Html:    body,
	}
	_, err := m.client.Emails.Send(params)
	if err != nil {
		m.log.Errorf("%s:%v", m.ServiceName(), err)
		return err
	}
	return nil
}

func (m MailingServices) ServiceName() string {
	return "Mailing"
}
