package stringutils

var alphabetSize = 256

// Adopted from https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
func DamerauLevenshteinDistance(sourceString, targetString string) (distance int) {
	letterFrequencyCache := make([]int, alphabetSize)
	for i := range letterFrequencyCache {
		letterFrequencyCache[i] = 0
	}

	maxdist := len(sourceString) + len(targetString)
	distanceTracker := make([][]int, len(sourceString)+2)
	for i := range distanceTracker {
		distanceTracker[i] = make([]int, len(targetString)+2)
		if i < 2 {
			continue
		}
		distanceTracker[i][0] = maxdist
		distanceTracker[i][1] = i
	}

	for j := 2; j < len(targetString)+2; j++ {
		distanceTracker[0][j] = maxdist
		distanceTracker[1][j] = j
	}

	for i := 2; i < len(sourceString)+2; i++ {
		lastMatchedTargetCharIdx := 0
		for j := 2; j < len(targetString)+2; j++ {
			var cost int
			prevSrcMatchIdx := letterFrequencyCache[targetString[j-2]]
			prevTgtMatchIdx := lastMatchedTargetCharIdx
			if sourceString[i-2] == targetString[j-2] {
				cost = 0
				lastMatchedTargetCharIdx = j
			} else {
				cost = 1
			}

			distanceTracker[i][j] = min(
				distanceTracker[i-1][j-1]+cost, // substitution
				distanceTracker[i][j-1]+1,      // insertion
				distanceTracker[i-1][j]+1,      // deletion
				distanceTracker[prevSrcMatchIdx+1][prevTgtMatchIdx+1]+(i-prevSrcMatchIdx-2)+1+(j-prevTgtMatchIdx-2), // transposition
			)
		}
	}

	return distanceTracker[len(sourceString)+1][len(targetString)+1]
}
