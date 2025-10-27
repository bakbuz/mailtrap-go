package common

type MediaType string

const (
	MediaType_ApplicationJSON MediaType = "application/json"
	MediaType_ApplicationXML  MediaType = "application/xml"
	MediaType_ImageJPEG       MediaType = "image/jpeg"
	MediaType_ImagePNG        MediaType = "image/png"
	MediaType_TextPlain       MediaType = "text/plain"
	MediaType_TextHTML        MediaType = "text/html"
)
