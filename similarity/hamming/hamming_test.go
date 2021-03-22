package hamming

import (
	"fmt"
	"testing"
)

func TestHammingDistance(test *testing.T) {

	var tests = []struct {
		s1, s2   string
		expected int
	}{{"", "", 0},
		{"A", "A", 0},
		{"ABC", "ABC", 0},
		{"ABC", "ABD", 1},
		{"ABC", "AB", -1},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("Hamming:%s with %s ... \n", tt.s1, tt.s2)

		test.Run(testname, func(t *testing.T) {
			result := Distance(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})

	}
}
