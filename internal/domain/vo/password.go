package vo

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/illusory-server/accounts/pkg/errors/codes"
	"github.com/illusory-server/accounts/pkg/errors/errx"
)

const (
	MinPasswordLen = 8
	MaxPasswordLen = 256
)

type Password struct {
	value string
}

func NewPassword(value string) (Password, error) {
	result := Password{value: value}

	if err := result.Validate(); err != nil {
		return Password{}, errx.WrapWithCode(err, codes.InvalidArgument, "Password.Validate")
	}

	return result, nil
}

func (p Password) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(
			&p.value,
			validation.Required,
			validation.Length(MinPasswordLen, MaxPasswordLen),
		),
	)
}

func (p Password) Value() string {
	return p.value
}

func (p Password) MarshalJSON() ([]byte, error) {
	return json.Marshal("secret-value")
}
