package gemini

// A ResponseWriter interface is used by a Gemini Protocol handler to construct a Gemini Protocol response.
//
// A ResponseWriter may not be used after the Handler.ServeGemini method has returned.
type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteHeader(statusCode int, mimeType string, mediaParameters map[string]string)
}
