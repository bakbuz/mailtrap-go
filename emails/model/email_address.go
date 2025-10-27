package model

// EmailAddress represents sender's or recipient's email address and name tuple, that can be used in From, To, CC or BCC parameters.
type EmailAddress struct {

	// Gets sender's or recipient's email address.
	//
	// Required. Must be valid email address.
	//
	// Contains sender's or recipient's email address.
	Email string `json:"email" validate:"required,email"`

	// Gets sender's or recipient's display name.
	//
	// Optional.
	//
	// Contains sender's or recipient's display name.
	DisplayName string `json:"name,omitempty"`
}

// Default instance constructor.
//
// Sender's or recipient's email address.
//
// Required. Must be valid email address.
//
// Optional sender's or recipient's display name.
func NewEmailAddress(email string, displayName string) *EmailAddress {

	if email == "" {
		panic("email cannot be empty")
	}

	return &EmailAddress{
		Email:       email,
		DisplayName: displayName,
	}
}
