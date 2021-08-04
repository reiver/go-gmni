package gemini

import (
	"errors"
)

var (
	errNilReceiver       = errors.New("gemini: nil receiver")
	errNilResponseWriter = errors.New("gemini: nil gemini.ResponseWriter")
	errNilWriter         = errors.New("gemini: nil io.Writer")
)
