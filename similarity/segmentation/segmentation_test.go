package segmentation

import (
	"fmt"
	"testing"
)

func TestSegmentation(t *testing.T) {
	var tests = []struct {
		s1, s2   string
		expected float64
	}{
		{"", "", 0.0},
		{"A", "A", 1.0},
		{"AB", "AB", 1.0},
		{"A B C", "A B C", 1.0},
		{"A B C", "A B D", 6.0 / 9.0},
		{"A B C", "A X C", 6.0 / 9.0},
		{"A B C", "A D E", 3.0 / 9.0},
		{"A B C", "A C B", 5.0 / 9.0},
		{"AB CD EF", "AB XX XX", 3.0 / 9.0},
		{"AB CD EF", "CD EF AB", 3.0 / 9.0},
		{"AB CD EF", "CD XY ZA", 1.0 / 9.0},
		{"AB CD EF", "CD AB XX", 2.0 / 9.0},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestSegmentation:'%s' & '%s' ... \n", tt.s1, tt.s2)

		t.Run(testname, func(t *testing.T) {
			s := CompareSegmentedStrings(tt.s1, tt.s2)

			if s != tt.expected {
				t.Errorf("got %v, want %v", s, tt.expected)
			}
		})

	}
}
