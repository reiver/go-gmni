package geminirequest

import (
	"github.com/reiver/go-utf8s"

	"io"
	"strings"
)

// Max represents the maximum number of bytes that a permitted, by the Gemini Protocol specification,
// for a Gemini Protocol request-line.
const Max = 1024

// Parse parses a Gemini Protocol request-line, and returns the target.
//
// For example, if the Gemini Protocol request-line was:
//
//	"gemini://example.com/path/to/something\r\n"
//
// The Parse would return:
//
//	"gemini://example.com/path/to/something"
//
// Gemini Protocol requests are a single CRLF-terminated line with the following structure:
//
// <TARGET><CR><LF>
//
// Where “<TARGET>” is a UTF-8 encoded string, with maximum length of 1024 bytes.
//
// Typically “<TARGET>” is an absolute URL, including the scheme.
//
// For example:
//
// gemini://example.com/path/to/something
//
// gemini://jam.strawberry/apple/banana/cherry/date
//
// gemini://peanut.butter/🙂.txt
//
// http://changelog.ca/
//
// hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8
//
// data:text/plain;charset=utf-8;base64,VGhpcyBpcyBhIHRlc3Qh
//
// Note that the when the target is a URL, it does NOT have to be a gemini-URL.
//
// It is up to a Gemini Protocol server to decide which target URLs it wants to respond to with a `20 SUCCESS` or not.
// This function just tries to parse it, and returns it.
func Parse(reader io.Reader) (string, error) {

	var target strings.Builder
	{
		var numIterations int
		var numBytes int

		for {
			numIterations++
			if Max < numIterations {
/////////////////////////////// RETURN
				return "", errBadRequest
			}

			r, size, err := utf8s.ReadRune(reader)
			numBytes += size
			if utf8s.RuneError == r {
/////////////////////////////// RETURN
				return "", errBadRequest
			}
			if nil != err && io.EOF != err {
/////////////////////////////// RETURN
				return "", err
			}

			if numBytes > Max {
/////////////////////////////// RETURN
				return "", errBadRequest
			}

			if '\r' == r {

				if io.EOF == err {
/////////////////////////////////////// RETURN
					return "", errBadRequest
				}

				r, size, err := utf8s.ReadRune(reader)
				numBytes += size
				if utf8s.RuneError == r {
/////////////////////////////////////// RETURN
					return "", errBadRequest
				}
				if nil != err && io.EOF != err {
/////////////////////////////////////// RETURN
					return "", err
				}

				if '\n' != r {
/////////////////////////////////////// RETURN
					return "",  errBadRequest
				}

		/////////////// BREAK
				break
			}

			target.WriteRune(r)

			if io.EOF == err {
/////////////////////////////// RETURN
				return "", errBadRequest
			}
		}
	}

	return target.String(), nil
}
