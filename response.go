package gemini

import (
	"io"
)

// Response represents the response from a Gemini Protocol request.
type Response struct {
	StatusCode int // ex: 20

	Meta string // ex: "text/gemini; lang=en,fa"

	Body io.ReadCloser
}
