package gemini

import (
	"strings"

	"testing"
)

func TestResponseWriter_writeHeader(t *testing.T) {

	tests := []struct {
		StatusCode int
		Meta       string

		Expected string
	}{
		{
			StatusCode: 10,
			Meta: "How many lights are there?",

			Expected: "10 How many lights are there?\r\n",
		},
		{
			StatusCode: 10,
			Meta: "Do you like rootbeer?",

			Expected: "10 Do you like rootbeer?\r\n",
		},
		{
			StatusCode: 10,
			Meta: "🙂",

			Expected: "10 🙂\r\n",
		},



		{
			StatusCode: 11,
			Meta: "Please choose a password",

			Expected: "11 Please choose a password\r\n",
		},
		{
			StatusCode: 11,
			Meta: "Tell me a secret",

			Expected: "11 Tell me a secret\r\n",
		},



		{
			StatusCode: 20,
			Meta: "text/gemini",

			Expected: "20 text/gemini\r\n",
		},
		{
			StatusCode: 20,
			Meta: "text/gemini; lang=en",

			Expected: "20 text/gemini; lang=en\r\n",
		},
		{
			StatusCode: 20,
			Meta: "text/gemini; lang=fa",

			Expected: "20 text/gemini; lang=fa\r\n",
		},
		{
			StatusCode: 20,
			Meta: "text/gemini; lang=en,fa",

			Expected: "20 text/gemini; lang=en,fa\r\n",
		},



		{
			StatusCode: 30,
			Meta: "gemini://example.com/path/to/file",

			Expected: "30 gemini://example.com/path/to/file\r\n",
		},
		{
			StatusCode: 30,
			Meta: "https://example.com/once/twice/thrice/fource",

			Expected: "30 https://example.com/once/twice/thrice/fource\r\n",
		},



		{
			StatusCode: 31,
			Meta: "gemini://example.com/path/to/file",

			Expected: "31 gemini://example.com/path/to/file\r\n",
		},
		{
			StatusCode: 31,
			Meta: "https://example.com/once/twice/thrice/fource",

			Expected: "31 https://example.com/once/twice/thrice/fource\r\n",
		},


		{
			StatusCode: 40,
			Meta: "crap! 💩",

			Expected: "40 crap! 💩\r\n",
		},
	}

	for testNumber, test := range tests {

		var buffer strings.Builder

		err := writeHeader(&buffer, test.StatusCode, test.Meta)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Log("ERROR:", err)
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d, the actual written header line was not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}

}
