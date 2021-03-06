package geminierror

import (
	"github.com/reiver/go-gemini/response/code"
	"github.com/reiver/go-gemini/response/name"

	"fmt"
	"strings"
)

type BadRequest interface {
	Error
	GeminiResponseBadRequest()
}

func BadRequestError(a ...interface{}) error {
	msg := fmt.Sprint(a...)

	return internalBadRequest{msg:msg}
}

func BadRequestErrorf(format string, a ...interface{}) error {
	msg := fmt.Sprintf(format, a...)

	return internalBadRequest{msg:msg}
}

type internalBadRequest struct {
	msg string
}

func (internalBadRequest) GeminiResponseBadRequest() {
	// Nothing here.
}

func (internalBadRequest) GeminiReponseStatusCode() int {
	return geminiresponsecode.BadRequest
}

func (internalBadRequest) GeminiReponseStatusName() string {
	return geminiresponsename.BadRequest
}

func (receiver internalBadRequest) Error() string {
	var buffer strings.Builder

	buffer.WriteString("gemini: response: ")
	fmt.Fprintf(&buffer, "%d %s", receiver.GeminiReponseStatusCode(), receiver.GeminiReponseStatusName())
	if msg := receiver.msg; "" != msg {
		fmt.Fprintf(&buffer, ": %s", msg)
	}

	return buffer.String()
}

// This will throw a compile-time error if the struct doesn't fit the interface.
// This saves us from having to write a unit test for this.
var _ BadRequest = internalBadRequest{}
