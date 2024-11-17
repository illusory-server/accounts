package vo

import (
	"encoding/json"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ID struct {
	value string
}

func NewID(v string) (ID, error) {
	result := ID{value: v}
	if err := result.Validate(); err != nil {
		return ID{}, err
	}

	return result, nil
}

func (v ID) Value() string {
	return v.value
}

func (v ID) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.value,
			validation.Required.Error("value is empty"),
			is.UUIDv4),
	)
}

func (v ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}
