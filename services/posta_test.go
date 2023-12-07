package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Posta_Services(t *testing.T) {
	t.Run("should_return_service_name", func(t *testing.T) {
		assert.Equal(t, postaService.ServiceName(), "postaClient")
	})

	t.Run("should_return_list_of_towns", func(t *testing.T) {
		towns, err := postaService.GetTowns()

		assert.Nil(t, err)
		assert.Equal(t, len(towns), 898)
	})

	t.Run("search_town_by_name", func(t *testing.T) {
		foundTowns, err := postaService.SearchTown("ngong hills")

		assert.Nil(t, err)
		assert.Equal(t, len(foundTowns), 1)
		assert.Equal(t, foundTowns[0].Town, "Ngong Hills")
		assert.Equal(t, foundTowns[0].PostalCode, "00208")
	})

	t.Run("search_non_existent_town", func(t *testing.T) {
		foundTowns, err := postaService.SearchTown("erkhls")

		assert.Nil(t, err)
		assert.Equal(t, len(foundTowns), 0)
	})
}
