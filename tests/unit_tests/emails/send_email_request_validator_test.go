package emails_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/maydere/mailtrap-go/emails/model"
	"gitlab.com/maydere/mailtrap-go/emails/request"
)

const (
	_validEmail   string = "someone@domean.com"
	_invalidEmail string = "someone"
	_templateId   string = "ID"
)

// --------------------------------------------------------------------------
// Request
// --------------------------------------------------------------------------

func TestValidation_ShouldFail_WhenNoRecipientsPresent(t *testing.T) {
	req := request.CreateSendEmailRequest()

	result := req.Validate()
	assert.NotNil(t, result)
}
  
  func TestValidation_ShouldNotFail_WhenOnlyToRecipientsPresent(t *testing.T)  {
      req := request.CreateSendEmailRequest()
	  req.To = []model.EmailAddress{*model.NewEmailAddress(_validEmail,"")}
          

      //var result = SendEmailRequestValidator.Instance.TestValidate(request);
      //result.ShouldNotHaveValidationErrorFor(r => r);
  }

  
  func TestValidation_ShouldNotFail_WhenOnlyCcRecipientsPresent(t *testing.T)  {
      var request = SendEmailRequest
          .Create()
          .Cc(_validEmail);

      //var result = SendEmailRequestValidator.Instance.TestValidate(request);
      //result.ShouldNotHaveValidationErrorFor(r => r);
  }

  
  func TestValidation_ShouldNotFail_WhenOnlyBccRecipientsPresent(t *testing.T)  {
      var request = SendEmailRequest
          .Create()
          .Bcc(_validEmail);

      //var result = SendEmailRequestValidator.Instance.TestValidate(request);
      //result.ShouldNotHaveValidationErrorFor(r => r);
  }