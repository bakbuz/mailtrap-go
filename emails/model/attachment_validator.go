package model

import (
	"github.com/bakbuz/mailtrap-go/core/validation"
)

func (m *Attachment) Validate() error {
	return validation.Validate.Struct(m)
}
