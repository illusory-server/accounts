package vo

import (
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRole(t *testing.T) {
	t.Run("Should correct role with role type value", func(t *testing.T) {
		tt := []struct {
			role vo.AccountRoleType
		}{
			{vo.RoleSuperAdmin},
			{vo.RoleAdmin},
			{vo.RoleUser},
		}

		for _, tc := range tt {
			role, err := vo.NewRole(tc.role)
			assert.NoError(t, err)
			assert.Equal(t, tc.role, role.Value())
		}
	})

	t.Run("Should error with incorrect role value", func(t *testing.T) {
		var incorrectRole vo.AccountRoleType = "KEK"
		role, err := vo.NewRole(incorrectRole)
		assert.Error(t, err)
		assert.Equal(t, vo.Role{}, role)
	})
}
