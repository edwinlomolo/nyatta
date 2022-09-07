package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	service "github.com/3dw1nM0535/nyatta/services"
)

type Resolver struct {
	UserService *service.UserServices
}
