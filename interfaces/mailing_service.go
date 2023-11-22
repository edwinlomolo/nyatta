package interfaces

import (
	"context"

	"github.com/3dw1nM0535/nyatta/graph/model"
)

type Mailing interface {
	ServiceName() string
	SaveMailing(ctx context.Context, email string) (*model.Status, error)
	SendEmail(ctx context.Context, to []string, from, subject, body string) error
}
