package segmentation

import (
	"strings"

	mmath "../minimath"
)

const (
	weightingExactMatch     = 0
	weightingExistanceMatch = 1
	weightingNoMatch        = 2
)

var segmentWeighting [weightingNoMatch + 1]float64

func contains(a []string, x string) int {
	for pos, n := range a {

		if x == n {

			return pos
		}
	}
	return -1
}

func setSegmentWeighting(longestStringLength int) {

	segmentWeighting[weightingExactMatch] = float64(longestStringLength)
	segmentWeighting[weightingExistanceMatch] = 1.0
	segmentWeighting[weightingNoMatch] = 0.0

}

//just break string up on whitespace
func segmentString(s string) []string {
	return strings.Fields(s)
}

// CompareSegmentedStrings - determines a measure of how similar two strings are based on segment position
func CompareSegmentedStrings(s1, s2 string) float64 {
	similarity := 0.0

	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	if s1 == s2 {
		return 1
	}

	segs1 := segmentString(s1) // target string comparing against
	segs2 := segmentString(s2)

	longestStringLength := mmath.Max2(len(segs1), len(segs2))
	setSegmentWeighting(longestStringLength)

	// need to have something to compare
	if len(segs1) == 0 || len(segs2) == 0 {
		return similarity
	}

	if len(segs1) > len(segs2) {
		for index, segment := range segs2 {
			pos := contains(segs1, segment)

			if pos == index {
				// same segment in same place

				similarity += segmentWeighting[weightingExactMatch]
			} else if pos != -1 {
				similarity += segmentWeighting[weightingExistanceMatch]
			}
		}
	} else {
		for index, segment := range segs1 {
			pos := contains(segs2, segment)

			if pos == index {
				// same segment in same place
				similarity += segmentWeighting[weightingExactMatch]
			} else if pos != -1 {
				similarity += segmentWeighting[weightingExistanceMatch]
			}
		}
	}

	return float64(similarity / (segmentWeighting[weightingExactMatch] * float64(longestStringLength)))
}
