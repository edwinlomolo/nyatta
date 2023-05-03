package resolver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Postal_Resolver(t *testing.T) {
	srv := makeAuthedGqlServer(true, ctx)

	t.Run("should_query_town_details", func(t *testing.T) {
		var searchTown struct {
			SearchTown []struct {
				ID         string
				Town       string
				PostalCode string
			}
		}

		query := fmt.Sprintf(
			`query { searchTown(town: %q) { id, town, postalCode } }`,
			"Ngong Hills",
		)

		srv.MustPost(query, &searchTown)

		assert.Equal(t, len(searchTown.SearchTown), 1)
		assert.Equal(t, searchTown.SearchTown[0].Town, "Ngong Hills")
		assert.Equal(t, searchTown.SearchTown[0].PostalCode, "00208")
	})
}
