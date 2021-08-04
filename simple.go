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
func RedirectTemporary(w ResponseWriter, target interface{}) error {
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
func RedirectPermanent(w ResponseWriter, target interface{}) error {
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

// TemporaryFailure writes a `40 TEMPORARY FAILURE` Gemini Protocol response to `w`.
func TemporaryFailure(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(message)

	const statusCode = 40

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// ServerUnavailable writes a `41 SERVER UNAVAILABLE` Gemini Protocol response to `w`.
func ServerUnavailable(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(message)

	const statusCode = 41

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// CGIError writes a `42 CGI ERROR` Gemini Protocol response to `w`.
func CGIError(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(message)

	const statusCode = 42

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// ProxyError writes a `43 PROXY ERROR` Gemini Protocol response to `w`.
func ProxyError(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(message)

	const statusCode = 43

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// SlowDown writes a `44 SLOW DOWN` Gemini Protocol response to `w`.
func SlowDown(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(message)

	const statusCode = 44

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// PermanentFailure writes a `50 PERMANENT FAILURE` Gemini Protocol response to `w`.
func PermanentFailure(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(message)

	const statusCode = 50

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

// Gone writes a `52 GONE` Gemini Protocol response to `w`.
func Gone(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(message)

	const statusCode = 52

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// ProxyRequestRefused writes a `53 PROXY REQUEST REFUSED` Gemini Protocol response to `w`.
func ProxyRequestRefused(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	var meta string = caststring(message)

	const statusCode = 53

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

// ClientCertificateRequired writes a `60 CLIENT CERTIFICATE REQUIRED` Gemini Protocol response to `w`.
func ClientCertificateRequired(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	const statusCode = 60

	var meta string = caststring(message)

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// CertificateNotAuthorised writes a `61 CERTIFICATE NOT AUTHORISED` Gemini Protocol response to `w`.
func CertificateNotAuthorised(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	const statusCode = 61

	var meta string = caststring(message)

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}

// CertificateNotValid writes a `62 CERTIFICATE NOT VALID` Gemini Protocol response to `w`.
func CertificateNotValid(w ResponseWriter, message interface{}) error {
	if nil == w {
		return errNilResponseWriter
	}

	const statusCode = 62

	var meta string = caststring(message)

	if err := w.WriteHeader(statusCode, meta); nil != err {
		return err
	}

	return nil
}
