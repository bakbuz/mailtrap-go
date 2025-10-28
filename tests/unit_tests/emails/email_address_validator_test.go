package emails_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/bakbuz/mailtrap-go/emails/model"
)

func TestValidation_ShouldFail_WhenProvidedEmailIsInvalid(t *testing.T) {
	recipient := &model.EmailAddress{Email: "abcdefg"}

	result := recipient.Validate()

	assert.NotNil(t, result)
}

func TestValidation_ShouldNotFail_WhenProvidedEmailIsValid(t *testing.T) {
	recipient := &model.EmailAddress{Email: "efe.kaya@domain.com"}

	result := recipient.Validate()

	assert.Nil(t, result)
}

func TestValidation_ShouldNotFail_WhenDisplayNameIsNotEmpty(t *testing.T) {
	recipient := &model.EmailAddress{Email: "ege.kaya@domain.com", DisplayName: "Ege Kaya"}

	result := recipient.Validate()

	assert.Nil(t, result)
}
