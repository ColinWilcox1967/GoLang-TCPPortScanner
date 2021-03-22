package similarity

import (

	// add any new algos here
	hamming "../similarity/hamming"
	jarowinkler "../similarity/jarowinkler"
	levenshtein "../similarity/levenshtein"
	segmentation "../similarity/segmentation"
	smithwaterman "../similarity/smithwaterman"

	"strings"

	fileUtils "../fileutils"
)

const (
	algorithmsFile string = "ALGORITHM.LIST" //??
)

func loadAlgorithmFile(fileName string) bool {
	_, err := fileUtils.OpenAndReadSettingsFile(fileName)
	if err != nil {
		return false
	}

	return true
}

func applyAlgorithm(settings []string, algorithmKey string) bool {
	for _, setting := range settings {
		if strings.Contains(setting, algorithmKey) {
			// found the key now extract whether we are going to use it
			parts := strings.Fields("=")

			// only use text to the right of the '='
			if strings.Contains(strings.ToUpper(parts[1]), "TRUE") {
				return true
			}

			return false
		}
	}

	return false
}

func maintainUpperAndLowerLimits(lowerLimit *float64, upperLimit *float64, similarity float64) {
	if similarity < *lowerLimit {
		*lowerLimit = similarity
	}

	if similarity > *upperLimit {
		*upperLimit = similarity
	}
}

// StringSimilarity Calculates string similarity as a percentage
func StringSimilarity(s1, s2 string) float64 {

	algorithmCount := 0    // number of alogrithms applied so far
	similarityTotal := 0.0 // running total of each similarity measure

	// maintain highest and lowest similarity measures
	mostSimilarMeasure := 0.0
	leastSimilarMeasure := 100.0

	if applyHammingDistance {
		similarity := float64(hamming.Distance(s1, s2))
		algorithmCount++
		maintainUpperAndLowerLimits(&leastSimilarMeasure, &mostSimilarMeasure, similarity)
		similarityTotal += similarity
	}

	if applyJaroWinkler {
		var similarity float64

		similarity = float64(jarowinkler.Distance(s1, s2))
		algorithmCount++
		maintainUpperAndLowerLimits(&leastSimilarMeasure, &mostSimilarMeasure, similarity)
		similarityTotal += similarity
	}

	if applyLevenshtein {
		similarity, _ := float64(levenshtein.Distance(s1, s2))
		algorithmCount++
		maintainUpperAndLowerLimits(&leastSimilarMeasure, &mostSimilarMeasure, similarity)
		similarityTotal += similarity
	}
	if applyNicknames {
		//	nicknames := nicknames.Distance(s1, s2)
		//algorithmCount++

		//maintainUpperAndLowerLimits (&leastSimilarMeasure, &mostSimilarMeasure, similarity)
		//similarityTotal += similarity
	}
	if applySegmentation {
		similarity := segmentation.CompareSegmentedStrings(s1, s2)
		algorithmCount++
		maintainUpperAndLowerLimits(&leastSimilarMeasure, &mostSimilarMeasure, similarity)
		similarityTotal += similarity
	}
	if applySmithWaterman {
		similarity := smithwaterman.Distance(s1, s2)
		algorithmCount++
		maintainUpperAndLowerLimits(&leastSimilarMeasure, &mostSimilarMeasure, similarity)
		similarityTotal += similarity
	}

	if algorithmCount == 0 { // only apply if weve applied at least one algorithm
		return 0.0
	}

	// Remove upper and lower edge values
	return (similarityTotal - mostSimilarMeasure - leastSimilarMeasure) / float64(algorithmCount)
}
