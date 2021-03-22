package hamming

//Distance - Calculates hamming distance between two strings of same length
func Distance(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}

	distance := 0
	for index := 0; index < len(s1); index++ {
		if s1[index] != s2[index] {
			distance++
		}
	}

	return distance
}
