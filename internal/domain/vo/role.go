package vo

import (
	validation "github.com/go-ozzo/ozzo-validation"
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
		return Role{}, err
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
