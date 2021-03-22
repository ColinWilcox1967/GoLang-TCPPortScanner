package levenshtein

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	var tests = []struct {
		s1, s2   string
		expected int
	}{
		{"GILLY", "GEELY", 20},
		{"HONDA", "HYUNDAI", 3},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestDistance:'%s' & '%s' ... \n", tt.s1, tt.s2)

		t.Run(testname, func(t *testing.T) {
			d := Distance(tt.s1, tt.s2)

			if d != tt.expected {
				t.Errorf("got %v want %v", d, tt.expected)
			}
		})

	}
}
