package second_task

func getSimilarityScore(firstList, secondList []int) int {
	// create lookup: number to the number of times it occurs in the list
	lookup := map[int]int{}
	for _, secondListItem := range secondList {
		lookup[secondListItem]++
	}

	// calculate the similarity score
	similarityScore := 0
	for _, firstListItem := range firstList {
		if numTimes, ok := lookup[firstListItem]; ok {
			similarityScore += numTimes * firstListItem
		}
	}

	return similarityScore
}
