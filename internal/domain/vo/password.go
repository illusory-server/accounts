package vo

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

const (
	MinPasswordLen = 8
	MaxPasswordLen = 64
)

var isHasSymbol = regexp.MustCompile("([A-Za-z])")
var isHasDigit = regexp.MustCompile("([0-9])")

type Password struct {
	value string
}

func NewPassword(value string) (Password, error) {
	result := Password{value: value}

	if err := result.Validate(); err != nil {
		return Password{}, err
	}

	return result, nil
}

func (p Password) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(
			&p.value,
			validation.Required,
			validation.Length(MinPasswordLen, MaxPasswordLen),
			validation.Match(isHasSymbol),
			validation.Match(isHasDigit),
		),
	)
}

func (p Password) Value() string {
	return p.value
}
