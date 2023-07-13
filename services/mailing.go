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
	newUserEmail     = "<p>Welcome to Nyatta! I'm excited to have you onboard!<br /><br /> I'm working hard to make Nyatta the best place to find and share your favorite places with friends. You may encounter or face some issues but I'm able to track them as they happen. However, if you see anything of concern, I'm happy to have a chat.<br /><br /> More updates get pushed out every week, so keep an eye out for new features and major improvements.<br /><br /> Please reach out if you have any questions or feedback for me. I'm looking to improve and make your experience using Nyatta better!<br /><br /> <strong>Regards,</strong><br /> Edwin Lomolo.</p>"
	newPropertyEmail = "<p>I've received your listing! Here are the next steps: <ul><li>Prepare your property unit for a professional photo shoot session with us</li><li>Start adding tenants to your units</li></ul> I'm always pushing product updates weekly and you will receive an email discussing about what has changed/what's new. I'm always looking into how I can make your experience using Nyatta better. Don't hesitate to reach out to me!<br /><br /> <strong>Regards,</strong><br /> Edwin Lomolo.</p>"
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
		Html:    body,
	}
	_, err := m.client.Emails.Send(params)
	if err != nil {
		m.log.Errorf("error sending email: %s:%v", m.ServiceName(), err)
		return err
	}
	return nil
}

func (m MailingServices) ServiceName() string {
	return "Mailing"
}
