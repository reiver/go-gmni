package gemini

import (
	"fmt"
	"mime"
)

// Input writes a `10 INPUT` Gemini Protocol response to `w`.
func Input(w ResponseWriter, prompt interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(prompt)

	const statusCode = 10

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// SensitiveInput writes an `11 SENSITIVE INPUT` Gemini Protocol response to `w`.
func SensitiveInput(w ResponseWriter, prompt interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(prompt)

	const statusCode = 11

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// Success writes a `20 SUCCESS` Gemini Protocol response to `w`.
func Success(w ResponseWriter, mediaType string, mediaParameters map[string]string, body []byte) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = mime.FormatMediaType(mediaType, mediaParameters)

	const statusCode = 20

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	{
		expected := len(body)

		n, err := w.Write(body)
		if nil != err {
			return err
		}
		if actual := n; expected != actual {
			return fmt.Errorf("gemini: problem writing Gemini Protocol body; expected to write %d bytes, but actually wrote %d bytes", expected, actual)
		}
	}

	return nil
}

// RedirectTemporary writes a `30 REDIRECT - TEMPORARY` Gemini Protocol response to `w`.
func RedirectTemporary(w ResponseWriter, target  interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(target)

	const statusCode = 30

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// RedirectPermanent writes a `31 REDIRECT - PERMANENT` Gemini Protocol response to `w`.
func RedirectPermanent(w ResponseWriter, target  interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(target)

	const statusCode = 31

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// NotFound writes a `51 NOT FOUND` Gemini Protocol response to `w`.
func NotFound(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	const statusCode = 51

	var meta string = caststring(message)

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// BadRequest writes a `59 BAD REQUEST` Gemini Protocol response to `w`.
func BadRequest(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	const statusCode = 59

	var meta string = caststring(message)

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}
