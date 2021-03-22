package jarowinkler

import (
	mmath "../minimath"
)

func contains(s string, b byte) int {
	for index := 0; index < len(s); index++ {
		if s[index] == b {
			return index
		}
	}
	return -1
}

func maxDistanceAway(s1, s2 string) int {

	distance := float64(mmath.Max2(len(s1), len(s2)) / 2)

	return int(mmath.Floor(distance)) - 1
}

func countTranspositions(s1, s2 string) int {

	matches := 0
	for index := 0; index < len(s1); index++ {

		if pos := contains(s2, s1[index]); pos != -1 {
			positionSeparation := mmath.Absolute(pos - index)

			if positionSeparation > 0 && positionSeparation < maxDistanceAway(s1, s2) {
				matches++
			}
		}

	}

	return matches
}

func countMatchingCharacterPairs(s1, s2 string) int {

	// only compare upto length of shortest string
	numberOfPairsToCompare := mmath.Min2(len(s1), len(s2))

	count := 0
	for index := 0; index < numberOfPairsToCompare; index++ {
		if s1[index] == s2[index] {
			count++
		}
	}
	return count

}

// Distance calculates Jaro-Winkler distance between two strings
func Distance(s1, s2 string) float64 {

	if len(s1) == 0 || len(s2) == 0 {
		return 0.0
	}

	if s1 == s2 {
		return 1.0
	}

	m := countMatchingCharacterPairs(s1, s2)
	t := countTranspositions(s1, s2)

	if m == 0 {
		return 0.0
	}

	retval := float64(m) / float64(len(s1))
	retval += float64(m) / float64(len(s2))
	retval += float64(float64(m-t) / float64(m))
	retval /= 3.0

	return retval
}
