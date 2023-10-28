package resolver

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func Test_Resolver_User(t *testing.T) {
	var user *model.User

	t.Run("resolver_should_get_user", func(t *testing.T) {
		var err error

		user, err = userService.FindUserByPhone("+254829639846")
		if err != nil {
			t.Errorf("expected nil err got %v", err)
		}
	})

	t.Run("resolver_should_get_properties_belonging_to_user", func(t *testing.T) {
		_, err = propertyService.CreateProperty(&model.NewProperty{
			Name:       "Ngong Hills Agency",
			PostalCode: "00208",
			Town:       "Ngong Hills",
			CreatedBy:  user.ID,
		})
		assert.Nil(t, err)
	})
}
