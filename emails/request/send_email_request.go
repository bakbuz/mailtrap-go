package request

import "gitlab.com/maydere/mailtrap-go/emails/model"

// Represents request object used to send email.
type SendEmailRequest struct {

	// Gets or sets EmailAddress instance representing email's sender.
	//
	// Required.
	//
	// Instance, representing email's sender address and name.
	From *model.EmailAddress `json:"from" validate:"required"`

	// Gets or sets EmailAddress representing 'Reply To' email field.
	//
	// Representing 'Reply To' address and name.
	ReplyTo *model.EmailAddress `json:"reply_to,omitempty"`

	// Gets a collection of EmailAddress objects, defining who will receive a copy of email.
	//
	// A collection of EmailAddress objects.
	//
	// Must not contain more than 1000 recipients.
	// Each object in this collection must contain the recipient's email address.
	// Each object in this collection may optionally contain the recipient's name.
	// At least one recipient must be specified in one of the collections: To, Cc or Bcc.
	To []model.EmailAddress `json:"to" validate:"required,min=1,max=1000,dive"`

	// Gets a collection of EmailAddress objects, defining who will receive a carbon copy of email.
	//
	// A collection of EmailAddress objects.
	//
	// Must not contain more than 1000 recipients.
	// Each object in this collection must contain the recipient's email address.
	// Each object in this collection may optionally contain the recipient's name.
	// At least one recipient must be specified in one of the collections: To, Cc or Bcc.
	Cc []model.EmailAddress `json:"cc,omitempty" validate:"omitempty,max=1000,dive"`

	// Gets a collection of EmailAddress objects, defining who will receive a blind carbon copy of email.
	//
	// A collection of EmailAddress objects.
	//
	// Must not contain more than 1000 recipients.
	// Each object in this collection must contain the recipient's email address.
	// Each object in this collection may optionally contain the recipient's name.
	// At least one recipient must be specified in one of the collections: To, Cc or Bcc.
	Bcc []model.EmailAddress `json:"bcc,omitempty" validate:"omitempty,max=1000,dive"`

	// Gets a collection of Attachment objects, where you can specify any attachments you want to include.
	//
	// A collection of Attachment objects.
	Attachments []model.Attachment `json:"attachments,omitempty"`

	// Gets a dictionary of header names and values to substitute for them.
	//
	// A dictionary of header names and values.
	//
	// The key/value pairs must be strings.
	// You must ensure these are properly encoded if they contain unicode characters.
	// These headers cannot be one of the reserved headers.
	// "Content-Transfer-Encoding" header will be ignored and replaced with "quoted-printable".
	Headers map[string]string `json:"headers,omitempty"`

	// Gets a dictionary of values that are specific to the entire send
	// that will be carried along with the email and its activity data.
	//
	// A dictionary of variable keys and values.
	//
	// The key/value pairs must be strings.
	// Total size of custom variables in JSON form must not exceed 1000 bytes.
	CustomVariables map[string]string `json:"custom_variables,omitempty"`

	// Gets or sets the global or 'message level' subject of email.
	// This may be overridden by subject lines set in personalizations.
	//
	// Contains the subject of the email.
	//
	// Must be null if TemplateId is set.
	//
	// Required in case HtmlBody and(or) TextBody is used.
	// Should be non-empty string in this case.
	Subject string `json:"subject,omitempty"`

	// Gets or sets the text version of the body of the email.
	//
	// Contains the text body of the email.
	//
	// Must be null if TemplateId is set.
	// Otherwise, can be used along with HtmlBody to create a fall-back for non-html clients.
	// Required in the absence of TemplateId and HtmlBody.
	TextBody string `json:"text,omitempty"`

	// Gets or sets HTML version of the body of the email.
	//
	// Contains the HTML body of the email.
	//
	// Must be null if TemplateId is set.
	// Required in the absence of TemplateId and TextBody.
	HtmlBody string `json:"html,omitempty"`

	// Gets or sets the category of email.
	//
	// Contains the category of the email.
	//
	// Should be null if TemplateId is set.
	// Otherwise must be less or equal to 255 characters.
	Category string `json:"category,omitempty" validate:"omitempty,max=255"`

	// Gets or sets UUID of email template.
	//
	// Contains the UUID of email template.
	//
	// If provided, then Subject, Category, TextBody  and HtmlBody
	// properties are forbidden and must be null.
	// Email subject, text and html will be generated from template using optional TemplateVariables.
	TemplateId string `json:"template_uuid,omitempty"`

	// Gets or sets optional template variables that will be used to generate actual subject, text and html
	// from email template.
	//
	// Contains template variables object.
	//
	// Will be used only in case TemplateId is set.
	TemplateVariables interface{} `json:"template_variables,omitempty"`
}

func CreateSendEmailRequest() *SendEmailRequest {
	return &SendEmailRequest{}
}
