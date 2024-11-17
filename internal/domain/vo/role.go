package vo

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/illusory-server/accounts/pkg/errors/codes"
	"github.com/illusory-server/accounts/pkg/errors/xerr"
)

type AccountRoleType string

const (
	RoleAdmin      AccountRoleType = "ADMIN"
	RoleSuperAdmin AccountRoleType = "SUPER_ADMIN"
	RoleUser       AccountRoleType = "USER"
)

type Role struct {
	value AccountRoleType
}

func NewRole(value AccountRoleType) (Role, error) {
	result := Role{value: value}

	if err := result.Validate(); err != nil {
		return Role{}, xerr.WrapWithCode(err, codes.Unprocessable, "Role.Validate")
	}

	return result, nil
}

func (r Role) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(
			&r.value,
			validation.Required,
			validation.In(RoleAdmin, RoleUser, RoleSuperAdmin),
		),
	)
}

func (r Role) Value() AccountRoleType {
	return r.value
}

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value)
}
