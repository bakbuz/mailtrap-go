package response

// Hata oluşturucuyu kullanmak için (NewErrorResponse içinde)

// Represents send email response object.
type SendEmailResponse struct {
	// Gets a flag, indicating whether request succeeded or failed and response contains error(s).
	//
	// False when request failed and response contains error(s).
	//
	// True when request succeeded.
	Success bool `json:"success"`

	// Gets errors, associated with the response.
	//
	// Collection of error(s) details.
	ErrorData []string `json:"errors,omitempty"`

	// Gets a collection of IDs of emails that have been sent.
	//
	// A collection of IDs of emails that have been sent.
	MessageIds []string `json:"message_ids,omitempty"`
}

// Default instance constructor.
//
// Succeeded flag, indicating whether request succeeded and response a collection of message IDs.
func NewSuccessResponse(messageIds []string) *SendEmailResponse {

	if len(messageIds) == 0 {
		messageIds = make([]string, 0)
	}

	return &SendEmailResponse{
		Success:    true,
		MessageIds: messageIds,
		ErrorData:  nil,
	}
}

// Default instance constructor.
//
// Failed flag, indicating whether request failed and response contains errors to associate with the response.
func NewErrorResponse(errorData []string) *SendEmailResponse {

	if len(errorData) == 0 {
		errorData = []string{"An unspecified error occurred."}
	}

	return &SendEmailResponse{
		Success:    false,
		ErrorData:  errorData,
		MessageIds: nil,
	}
}

// Gets an empty response object.
var EmptySendEmailResponse = NewErrorResponse([]string{"Empty response."})
