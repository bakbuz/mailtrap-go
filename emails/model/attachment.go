package model

import "gitlab.com/maydere/mailtrap-go/core/common"

// Represents email attachment.
type Attachment struct {
	// Gets the Base64 encoded content of the attachment.
	//
	// Required. Must be non-empty string.
	//
	// Contains Base64 encoded content of the attachment.
	Content string `json:"content" validate:"required"`

	// Gets attachment's file name.
	//
	// Required. Must be non-empty string.
	//
	// Contains attachment file name.
	FileName string `json:"filename" validate:"required"`

	// Gets the attachment's content disposition, specifying how the attachment will be displayed.
	//
	// Optional.
	//
	// Indicates attachment's content disposition.
	//
	// - Inline results in the attached file are displayed automatically within the message.
	// - Attachment results in the attached file require some action to be taken before it is displayed,
	//  such as opening or downloading the file.
	//
	// Defaults to Attachment, if not specified explicitly.
	Disposition common.DispositionType `json:"disposition,omitempty"`

	// Gets the attachment's MIME type identifier.
	//
	// Optional.
	//
	// Contains the attachment's MIME type identifier.
	// E.g. "text/plain", "application/pdf", etc.
	MimeType common.MediaType `json:"type,omitempty" validate:"omitempty,min=3"`

	// Gets the attachment's content ID.
	//
	// Optional.
	//
	// Contains the attachment's content ID.
	//
	// This is used when the disposition is set to DispositionType.Inline and the attachment is an image,
	// allowing the file to be displayed within the body of your email.
	ContentId string `json:"content_id,omitempty"`
}

// Default instance constructor.
//
// content:
// The Base64 encoded content of the attachment.
// Required. Must be non-empty string.
//
//
// fileName:
// Attachment file name.
// Required. Must be non-empty string.
//
//
// disposition:
// The attachment's content disposition, specifying how you would like the attachment to be displayed.
// Optional. Defaults to DispositionType.Attachment if not specified explicitly.
//
//
// mimeType:
// Attachment MIME type identifier (e.g. "text/plain", "application/pdf", etc.)
// Optional.
//
//
// contentId:
// The attachment's content ID.
// Optional.
//
// panic:
//
// When content is null or empty.
//
// When fileName is null or empty.
func NewAttachment(content string, fileName string, disposition common.DispositionType, mimeType common.MediaType, contentId string) *Attachment {

	if content == "" {
		panic("content cannot be empty")
	}

	if fileName == "" {
		panic("fileName cannot be empty")
	}

	// Set default disposition if not provided
	if disposition == "" {
		disposition = common.DispositionType_Attachment
	}

	return &Attachment{
		Content:     content,
		FileName:    fileName,
		Disposition: disposition,
		MimeType:    mimeType,
		ContentId:   contentId,
	}
}
