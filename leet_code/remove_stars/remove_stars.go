package remove_stars

func removeStars(input string) string {
	//assumes we always have enough characters left of stars
	resultIndex := 0
	starCounter := 0
	result := make([]rune, len(input))
	for _, c := range input {
		if c == '*' {
			starCounter++
			continue
		} else {
			resultIndex -= starCounter
			result[resultIndex] = c
			resultIndex++
			starCounter = 0
		}
	}
	if starCounter > 0 {
		resultIndex -= starCounter
	}
	return string(result[:resultIndex])
}
