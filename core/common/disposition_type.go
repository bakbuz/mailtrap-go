package common

// Represents attachment disposition type.
type DispositionType string

const (
	// Gets disposition type that results in the attached file displayed automatically within the message.
	//
	// Static instance that results in the attached file displayed automatically within the message.
	DispositionType_Inline DispositionType = "inline"

	// Gets disposition type that results in the attached file require some action to be taken before it is displayed,
	// such as opening or downloading the file.
	//
	// Static instance that results in the attached file require some action to be taken before it is displayed.
	DispositionType_Attachment DispositionType = "attachment"
)
