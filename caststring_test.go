package gemini

import (
	"testing"
)

func TestCaststring(t *testing.T) {

	tests := []struct{
		Value interface{}
		Expected string
	}{
		{
			Value:    "",
			Expected: "",
		},
		{
			Value:    " ",
			Expected: " ",
		},
		{
			Value:    "apple BANANA Cherry",
			Expected: "apple BANANA Cherry",
		},
		{
			Value:    "Hello world! 🙂",
			Expected: "Hello world! 🙂",
		},



		{
			Value: []byte(""),
			Expected:     "",
		},
		{
			Value: []byte(" "),
			Expected:     " ",
		},
		{
			Value: []byte("apple BANANA Cherry"),
			Expected:     "apple BANANA Cherry",
		},
		{
			Value: []byte("Hello world! 🙂"),
			Expected:     "Hello world! 🙂",
		},



		{
			Value: []rune{},
			Expected:     "",
		},
		{
			Value: []rune{' '},
			Expected:     " ",
		},
		{
			Value: []rune{'a','p','p','l','e',' ','B','A','N','A','N','A',' ','C','h','e','r','r','y'},
			Expected:     "apple BANANA Cherry",
		},
		{
			Value: []rune{'H','e','l','l','o',' ','w','o','r','l','d','!',' ','🙂'},
			Expected:     "Hello world! 🙂",
		},

		{
			Value: rune(' '),
			Expected:   " ",
		},
		{
			Value: rune('a'),
			Expected:   "a",
		},
		{
			Value: rune('🙂'),
			Expected:   "🙂",
		},



		{
			Value: struct{Apple string; Banana int; Cherry bool}{Apple:"one", Banana:2, Cherry:true},
			Expected:   "{one 2 true}",
		},
	}

	for testNumber, test := range tests {

		actual := caststring(test.Value)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual result did not match what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
