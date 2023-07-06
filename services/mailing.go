package services

import (
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
)

type MailingServices struct {
	queries *sqlStore.Queries
}

// Enforce Mailing service interface
var _ interfaces.Mailing = &MailingServices{}

func NewMailingService(queries *sqlStore.Queries) *MailingServices {
	return &MailingServices{queries: queries}
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

func (m MailingServices) ServiceName() string {
	return "Mailing"
}
