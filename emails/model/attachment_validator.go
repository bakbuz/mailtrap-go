package model

import (
	"gitlab.com/maydere/mailtrap-go/core/validation"
)

func (m *Attachment) Validate() error {
	return validation.Validate.Struct(m)
}
