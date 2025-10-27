package request

import "github.com/go-playground/validator/v10"

// func (req *SendEmailRequest) Validate() *validation.ValidationResult {
// 	result := validation.NewValidationResult()

// 	if req.From == nil {
// 		result.AddError("From", "sender 'From' email address is required")
// 	} else {
// 		if err := req.From.Validate(); err != nil {
// 			result.AddError("From", err.Error())
// 		}
// 	}

// 	return result
// }

/*
func (req *SendEmailRequest) Validate() error {
	return validation.Validate.Struct(req)
}
*/

func (req *SendEmailRequest) Validate() error {
	return validate.Struct(req)
}

var validate = validator.New()

// Struct için özel validasyon kaydı
func init() {
	validate.RegisterStructValidation(SendEmailRequestValidation, SendEmailRequest{})
}

// Şartlı validasyon fonksiyonu:
func SendEmailRequestValidation(sl validator.StructLevel) {
	req := sl.Current().Interface().(SendEmailRequest)

	// Eğer TemplateId yoksa Subject zorunlu
	if req.TemplateId == "" {
		if req.Subject == "" {
			sl.ReportError(req.Subject, "Subject", "subject", "subject required if no template", "")
		}

		// TemplateId yoksa ve Subject varsa, body'lardan en az biri zorunlu
		if req.TextBody == "" && req.HtmlBody == "" {
			sl.ReportError(req.TextBody, "TextBody", "text", "if TemplateId is not present, at least one of the HtmlBody or TextBody fields is required. ", "")
			sl.ReportError(req.HtmlBody, "HtmlBody", "html", "if TemplateId is not present, at least one of the HtmlBody or TextBody fields is required.", "")
		}
	}
}
