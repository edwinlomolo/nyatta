package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
)

type Mailing interface {
	ServiceName() string
	SaveMailing(email string) (*model.Status, error)
}
