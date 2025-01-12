package vo

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Link struct {
	value string
}

func NewLink(value string) (Link, error) {
	res := Link{
		value: value,
	}
	if err := res.Validate(); err != nil {
		return Link{}, err
	}
	return res, nil
}

func (l Link) Value() string {
	return l.value
}

func (l Link) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.value, validation.Required, is.URL),
	)
}

func (l Link) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.value)
}
