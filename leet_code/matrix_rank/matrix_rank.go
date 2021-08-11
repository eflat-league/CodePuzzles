package matrix_rank

import (
	"sort"
)

func matrixRankTransform(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}

	//a map of each value to its position(s)
	positions := map[int][][2]int{}
	//a slice of all values
	nRows := len(matrix)
	nCols := len(matrix[0])
	allValuesSorted := make([]int, nRows*nCols)

	for i, row := range matrix {
		for j, value := range row {
			positions[value] = append(positions[value], [2]int{i, j})
			allValuesSorted[i*nCols+j] = value
		}
	}

	sort.Slice(allValuesSorted, func(i, j int) bool {
		return allValuesSorted[i] < allValuesSorted[j]
	})

	//keep track of max rank in each column and row
	//initialized with 0s
	rowMaxRank := make([]int, nRows)
	colMaxRank := make([]int, nCols)

	//initialize the result slice
	ranked := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		ranked[i] = make([]int, nCols)
	}

	for _, value := range deduplicate(allValuesSorted) {
		//since matrix can have repeating integers, we go through all position for any repeating value
		for _, pos := range positions[value] {
			//TODO solve rowMaxRank[pos[0]]/colMaxRank[pos[1]] for equal values that are in same row/col
			existingRank := max(rowMaxRank[pos[0]], colMaxRank[pos[1]])
			ranked[pos[0]][pos[1]] = existingRank + 1
		}
	}

	return ranked
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

//deduplicates a sorted slice;
func deduplicate(s []int) []int {
	var seen *int
	result := []int{}
	for _, value := range s {
		if seen == nil {
			seen = new(int)
			*seen = value
			result = append(result, value)
		} else {
			if *seen == value {
				continue
			}
			result = append(result, value)
			*seen = value
		}
	}
	return result
}

//given a slice, assigns unique rank to each item; NOT used in the solution of matrix rank;
func rankSlice(input []int) []int {
	sorted := make([]int, len(input))
	for i, value := range input {
		sorted[i] = value
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	rank := map[int]int{}
	for i, value := range sorted {
		rank[value] = i + 1 //we do not want 0 to be the starting rank, but 1
	}

	result := make([]int, len(input))
	for i, value := range input {
		result[i] = rank[value]
	}

	return result
}
