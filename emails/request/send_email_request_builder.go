package request

import (
	"github.com/bakbuz/mailtrap-go/core/common"
	"github.com/bakbuz/mailtrap-go/emails/model"
)

func (r *SendEmailRequest) WithFrom(email, name string) *SendEmailRequest {
	r.From = &model.EmailAddress{Email: email, DisplayName: name}
	return r
}

func (r *SendEmailRequest) WithReplyTo(email, name string) *SendEmailRequest {
	r.ReplyTo = &model.EmailAddress{Email: email, DisplayName: name}
	return r
}

func (r *SendEmailRequest) WithTo(email, name string) *SendEmailRequest {
	r.To = append(r.To, &model.EmailAddress{Email: email, DisplayName: name})
	return r
}

func (r *SendEmailRequest) WithCc(email, name string) *SendEmailRequest {
	r.Cc = append(r.Cc, &model.EmailAddress{Email: email, DisplayName: name})
	return r
}

func (r *SendEmailRequest) WithBcc(email, name string) *SendEmailRequest {
	r.Bcc = append(r.Bcc, &model.EmailAddress{Email: email, DisplayName: name})
	return r
}

func (r *SendEmailRequest) WithAttach(content, fileName string) *SendEmailRequest {
	a := &model.Attachment{
		Content:     content,
		FileName:    fileName,
		Disposition: common.DispositionType_Attachment,
	}
	r.Attachments = append(r.Attachments, a)
	return r
}

func (r *SendEmailRequest) WithAttach2(content, fileName string, disposition common.DispositionType, mimeType common.MimeType, contentId string) *SendEmailRequest {
	a := &model.Attachment{
		Content:     content,
		FileName:    fileName,
		Disposition: disposition,
		MimeType:    mimeType,
		ContentId:   contentId,
	}
	r.Attachments = append(r.Attachments, a)
	return r
}

func (r *SendEmailRequest) WithHeader(key, value string) *SendEmailRequest {
	r.Headers[key] = value
	return r
}

func (r *SendEmailRequest) WithCustomVariable(key, value string) *SendEmailRequest {
	r.CustomVariables[key] = value
	return r
}

func (r *SendEmailRequest) WithSubject(subject string) *SendEmailRequest {
	r.Subject = subject
	return r
}

func (r *SendEmailRequest) WithTextBody(text string) *SendEmailRequest {
	r.TextBody = text
	return r
}

func (r *SendEmailRequest) WithHtmlBody(html string) *SendEmailRequest {
	r.HtmlBody = html
	return r
}

func (r *SendEmailRequest) WithCategory(category string) *SendEmailRequest {
	r.Category = category
	return r
}

func (r *SendEmailRequest) WithTemplateId(id string) *SendEmailRequest {
	r.TemplateId = id
	return r
}

func (r *SendEmailRequest) WithTemplateVariables(vars map[string]interface{}) *SendEmailRequest {
	r.TemplateVariables = vars
	return r
}
