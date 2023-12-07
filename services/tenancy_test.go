package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tenancy_Services(t *testing.T) {
	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, tenancyService.ServiceName(), "tenancyClient")
	})
}
