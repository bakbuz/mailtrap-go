package model

import (
	"gitlab.com/maydere/mailtrap-go/core/validation"
)

func (m *EmailAddress) Validate() error {
	return validation.Validate.Struct(m)
}
