package gemini

import (
	"fmt"
	"io"
)

// A ResponseWriter interface is used by a Gemini Protocol handler to construct a Gemini Protocol response.
//
// A ResponseWriter may not be used after the Handler.ServeGemini method has returned.
type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteHeader(statusCode int, meta string) error
}

type internalResponseWriter struct {
	writer io.Writer
	headerWritten bool
}

func (receiver internalResponseWriter) Write(p []byte) (int, error) {

	writer := receiver.writer
	if nil == writer {
		return 0, errNilWriter
	}

	return writer.Write(p)
}

func (receiver *internalResponseWriter) WriteHeader(statusCode int, meta string) error {
	if nil == receiver {
		return errNilReceiver
	}

	writer := receiver.writer
	if nil == writer {
		return errNilWriter
	}

	err := writeHeader(writer, statusCode, meta)
	if nil != err {
		return err
	}

	receiver.headerWritten = true
	return nil
}

func writeHeader(writer io.Writer, statusCode int, meta string) error {

	if nil == writer {
		return errNilWriter
	}

	var header string = fmt.Sprintf("%d %s\r\n", statusCode, meta)

	var err error
	{
		var n int

		n, err = io.WriteString(writer, header)

		if nil != err {
			return err
		}

		if expected, actual := len(header), n; expected != actual {
			return fmt.Errorf("gemini: problem writing Gemini Protocol header %q; expected to write %d bytes, but actually wrote %d bytes", header, expected, actual)
		}
	}

	return nil
}
