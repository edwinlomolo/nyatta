package interfaces

import "github.com/3dw1nM0535/nyatta/graph/model"

type Posta interface {
	ServiceName() string
	GetTowns() ([]*model.Town, error)
	SearchTown(town string) ([]*model.Town, error)
}
