package emails_test

import (
	"testing"

	"github.com/bakbuz/mailtrap-go/emails/model"
	"github.com/bakbuz/mailtrap-go/emails/request"
	"github.com/stretchr/testify/assert"
)

const (
	_validEmail   string = "someone@domain.com"
	_invalidEmail string = "someone"
	_templateId   string = "ID"
)

// --------------------------------------------------------------------------
// Request
// --------------------------------------------------------------------------

func TestValidation_ShouldFail_WhenNoRecipientsPresent(t *testing.T) {
	req := request.Create()
	err := req.Validate()
	assert.NotNil(t, err)
}

func TestValidation_ShouldNotFail_WhenOnlyToRecipientsPresent(t *testing.T) {
	req := request.Create()
	req.From = model.NewEmailAddress(_validEmail, "")
	req.Subject = "subject"
	req.TextBody = "text body"

	req.To = []*model.EmailAddress{model.NewEmailAddress(_validEmail, "")}

	err := req.Validate()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)

	//var err = SendEmailRequestValidator.Instance.TestValidate(request);
	//err.ShouldNotHaveValidationErrorFor(r => r);
}

func TestValidation_ShouldNotFail_WhenOnlyCcRecipientsPresent(t *testing.T) {
	req := request.Create()
	req.From = model.NewEmailAddress(_validEmail, "")
	req.Subject = "subject"
	req.TextBody = "text body"

	req.Cc = []*model.EmailAddress{model.NewEmailAddress(_validEmail, "")}

	err := req.Validate()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
}

func TestValidation_ShouldNotFail_WhenOnlyBccRecipientsPresent(t *testing.T) {
	req := request.Create()
	req.From = model.NewEmailAddress(_validEmail, "")
	req.Subject = "subject"
	req.TextBody = "text body"

	req.Bcc = []*model.EmailAddress{model.NewEmailAddress(_validEmail, "")}

	err := req.Validate()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
}
