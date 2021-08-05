package geminirequest_test

import (
	"github.com/reiver/go-gemini/request"

	"strings"

	"testing"
)

func TestParse_success(t *testing.T) {

	tests := []struct{
		Input string
		Expected string
	}{
		{
			Input:    "\r\n",
			Expected: "",
		},



		{
			Input:    "gemini://example.com/\r\n",
			Expected: "gemini://example.com/",
		},
		{
			Input:    "gemini://example.com/path\r\n",
			Expected: "gemini://example.com/path",
		},
		{
			Input:    "gemini://example.com/path/to\r\n",
			Expected: "gemini://example.com/path/to",
		},
		{
			Input:    "gemini://example.com/path/to/something\r\n",
			Expected: "gemini://example.com/path/to/something",
		},
		{
			Input:    "gemini://example.com/path/to/something.txt\r\n",
			Expected: "gemini://example.com/path/to/something.txt",
		},



		{
			Input:    "gemini://jam.strawberry/\r\n",
			Expected: "gemini://jam.strawberry/",
		},
		{
			Input:    "gemini://jam.strawberry/jar.php\r\n",
			Expected: "gemini://jam.strawberry/jar.php",
		},
		{
			Input:    "gemini://jam.strawberry/apple/\r\n",
			Expected: "gemini://jam.strawberry/apple/",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana\r\n",
			Expected: "gemini://jam.strawberry/apple/banana",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana/\r\n",
			Expected: "gemini://jam.strawberry/apple/banana/",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana/cherry\r\n",
			Expected: "gemini://jam.strawberry/apple/banana/cherry",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana/cherry/\r\n",
			Expected: "gemini://jam.strawberry/apple/banana/cherry/",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana/cherry/date\r\n",
			Expected: "gemini://jam.strawberry/apple/banana/cherry/date",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana/cherry/date?a=1&bb=22&ccc=333\r\n",
			Expected: "gemini://jam.strawberry/apple/banana/cherry/date?a=1&bb=22&ccc=333",
		},



		{
			Input:    "gemini://peanut.butter/🙂.txt\r\n",
			Expected: "gemini://peanut.butter/🙂.txt",
		},
		{
			Input:    "gemini://peanut.butter/once/🙂.txt\r\n",
			Expected: "gemini://peanut.butter/once/🙂.txt",
		},
		{
			Input:    "gemini://peanut.butter/once/twice/🙂.txt\r\n",
			Expected: "gemini://peanut.butter/once/twice/🙂.txt",
		},
		{
			Input:    "gemini://peanut.butter/once/twice/thrice/🙂.txt\r\n",
			Expected: "gemini://peanut.butter/once/twice/thrice/🙂.txt",
		},
		{
			Input:    "gemini://peanut.butter/once/twice/thrice/fource/🙂.txt\r\n",
			Expected: "gemini://peanut.butter/once/twice/thrice/fource/🙂.txt",
		},



		{
			Input:    "http://changelog.ca/\r\n",
			Expected: "http://changelog.ca/",
		},



		{
			Input:    "hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8\r\n",
			Expected: "hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8",
		},



		{
			Input:    "data:text/plain;charset=utf-8;base64,VGhpcyBpcyBhIHRlc3Qh\r\n",
			Expected: "data:text/plain;charset=utf-8;base64,VGhpcyBpcyBhIHRlc3Qh",
		},



		{
			Input:    "apple banana cherry\r\n",
			Expected: "apple banana cherry",
		},
	}

	for testNumber, test := range tests {

		actual, err := geminirequest.Parse(strings.NewReader(test.Input))

		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Logf("INPUT:    %q", test.Input)
			t.Logf("EXPECTED: %q", test.Expected)
			t.Log("ERROR:", err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual result from parsing is not what was expected.", testNumber)
			t.Logf("INPUT:    %q", test.Input)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}

func TestParse_failure(t *testing.T) {

	tests := []struct{
		Input string
		Expected string
	}{
		{
			Input:    "",
			Expected: "gemini: response: 59 BAD REQUEST",
		},



		{
			Input:    "gemini://example.com/",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://example.com/path",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://example.com/path/to",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://example.com/path/to/something",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://example.com/path/to/something.txt",
			Expected: "gemini: response: 59 BAD REQUEST",
		},



		{
			Input:    "gemini://jam.strawberry/",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://jam.strawberry/jar.php",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://jam.strawberry/apple/",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana/",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana/cherry",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana/cherry/",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana/cherry/date",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://jam.strawberry/apple/banana/cherry/date?a=1&bb=22&ccc=333",
			Expected: "gemini: response: 59 BAD REQUEST",
		},



		{
			Input:    "gemini://peanut.butter/🙂.txt",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://peanut.butter/once/🙂.txt",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://peanut.butter/once/twice/🙂.txt",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://peanut.butter/once/twice/thrice/🙂.txt",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
		{
			Input:    "gemini://peanut.butter/once/twice/thrice/fource/🙂.txt",
			Expected: "gemini: response: 59 BAD REQUEST",
		},



		{
			Input:    "http://changelog.ca/",
			Expected: "gemini: response: 59 BAD REQUEST",
		},



		{
			Input:    "hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8",
			Expected: "gemini: response: 59 BAD REQUEST",
		},



		{
			Input:    "data:text/plain;charset=utf-8;base64,VGhpcyBpcyBhIHRlc3Qh",
			Expected: "gemini: response: 59 BAD REQUEST",
		},



		{
			Input:    "apple banana cherry",
			Expected: "gemini: response: 59 BAD REQUEST",
		},



		{
			Input:    "gemini://example.com/\rapple banana cherry date",
			Expected: "gemini: response: 59 BAD REQUEST",
		},
	}

	for testNumber, test := range tests {

		target, err := geminirequest.Parse(strings.NewReader(test.Input))

		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one.", testNumber)
			t.Logf("TARGET: %q", target)
			t.Logf("INPUT: %q", test.Input)
			t.Logf("EXPECTED: %q", test.Expected)
			continue
		}

		if expected, actual := test.Expected, err.Error(); expected != actual {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("TARGET: %q", target)
			t.Logf("INPUT:    %q", test.Input)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}

