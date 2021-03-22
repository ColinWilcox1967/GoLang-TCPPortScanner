package levenshtein

import mmath "../minimath"

// Distance between two strings
func Distance(s1, s2 string) int {

	if len(s1) == 0 || len(s2) == 0 {
		return 0.0
	}
	if s1 == s2 {
		return 1.0 // exact match
	}

	rows := len(s1) + 1
	columns := len(s2) + 1
	distanceMatrix := make([][]int, rows)
	for row := 1; row < rows; row++ {
		distanceMatrix[row] = make([]int, columns)
	}

	// Setup initial matrix
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			distanceMatrix[i][0] = i
			distanceMatrix[0][j] = j
		}
	}

	var cost int

	for column := 1; column < columns; column++ {
		for row := 1; row < rows; row++ {
			if s1[row-1] == s2[column-1] {
				cost = 0
			} else {
				cost = 2
			}

			distanceMatrix[row][column] = mmath.Min3(distanceMatrix[row-1][column]+1, // cost of deletion
				distanceMatrix[row][column-1]+1,      // cost of insertion
				distanceMatrix[row-1][column-1]+cost) // cost of substitution
		}

		//	ratio = (float64)(((len(s1) + len(s2)) - distanceMatrix[rows][columns]) / (len(s1) + len(s2)))
	}

	return distanceMatrix[rows][columns]
}
