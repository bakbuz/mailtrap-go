package model

import (
	"github.com/bakbuz/mailtrap-go/core/validation"
)

func (m *EmailAddress) Validate() error {
	return validation.Validate.Struct(m)
}
