package punctuation

import (
	"fmt"
	"testing"
)

func TestRemoveWhiteSpace(t *testing.T) {
	var tests = []struct {
		s              string
		expectedString string
	}{
		{"", ""},
		{" abc", "abc"},
		{"  abc", "abc"},
		{"abc ", "abc"},
		{"abc  ", "abc"},
		{"a b", "a b"},
		{"a  b", "a b"},
		{".ab", ".ab"},
		{"ab.", "ab."},
		{"..ab", ".ab"},
		{"ab..", "ab."},
		{"ab..cd", "ab.cd"},
		{"ab  cd  de", "ab cd de"},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestRemoveWhiteSpace:%s  ... \n", tt.s)

		t.Run(testname, func(t *testing.T) {
			s := RemovePunctuation(tt.s)

			if s != tt.expectedString {
				t.Errorf("got %v, want %v", s, tt.expectedString)
			}
		})

	}

}
