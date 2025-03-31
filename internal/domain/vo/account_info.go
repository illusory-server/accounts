package vo

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
)

const (
	MinNameLen = 2
	MaxNameLen = 124
)

type AccountInfo struct {
	firstName string
	lastName  string
	email     string
}

func NewAccountInfo(firstName, lastName, email string) (AccountInfo, error) {
	result := AccountInfo{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
	}

	if err := result.Validate(); err != nil {
		return AccountInfo{}, errx.WrapWithCode(err, codex.InvalidArgument, "AccountInfo.Validate")
	}

	return result, nil
}

func (a AccountInfo) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.firstName, validation.Required, validation.Length(MinNameLen, MaxNameLen)),
		validation.Field(&a.lastName, validation.Required, validation.Length(MinNameLen, MaxNameLen)),
		validation.Field(&a.email, validation.Required, is.Email),
	)
}

// getters

func (a AccountInfo) FirstName() string {
	return a.firstName
}
func (a AccountInfo) LastName() string {
	return a.lastName
}
func (a AccountInfo) Email() string {
	return a.email
}
func (a AccountInfo) FullName() string {
	return a.firstName + " " + a.lastName
}

// marshal

func (a AccountInfo) MarshalJSON() ([]byte, error) {
	marshal := map[string]interface{}{
		"first_name": a.firstName,
		"last_name":  a.lastName,
		"email":      a.email,
	}
	return json.Marshal(marshal)
}
