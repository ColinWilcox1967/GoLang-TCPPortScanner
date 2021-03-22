package jarowinkler

import (
	"fmt"
	"testing"

	minimath "../minimath"
)

const float64EqualityThreshold = float64(1e-9)

func almostEqual(f float64) bool {
	return float64(minimath.AbsoluteF(f)) <= float64EqualityThreshold
}

func TestJaroWinkler(test *testing.T) {
	var tests = []struct {
		s1, s2           string
		expectedDistance float64
	}{
		{"DANCE", "LANDS", float64(9.0 / 15.0)},
		{"CRATE", "TRACE", float64(11.0 / 15.0)},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("JaroWinkler:%s with %s ... \n", tt.s1, tt.s2)

		test.Run(testname, func(t *testing.T) {
			distance := Distance(tt.s1, tt.s2)

			if !almostEqual(distance - tt.expectedDistance) {
				t.Errorf("got %f, want %f", distance, tt.expectedDistance)
			}
		})

	}

}
