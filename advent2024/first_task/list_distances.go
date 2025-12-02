package first_task

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetDistance(firstList, secondList []int) int {
	sort.Slice(firstList, func(i, j int) bool {
		return firstList[i] < firstList[j]
	})
	sort.Slice(secondList, func(i, j int) bool {
		return secondList[i] < secondList[j]
	})

	if len(firstList) != len(secondList) {
		panic("lists must be same length")
	}

	totalDistance := 0
	for index, firstListItem := range firstList {
		totalDistance += absDiffInt(firstListItem, secondList[index])
	}

	return totalDistance
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// readInputFile reads the input file and creates two lists of integers.
func ReadInputFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list1 []int
	var list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // Split the line by whitespace

		if len(parts) == 2 { // Ensure there are two numbers on each line
			num1, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, err
			}
			list1 = append(list1, num1)

			num2, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, err
			}
			list2 = append(list2, num2)
		} //Could log lines that don't conform or return error
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return list1, list2, nil
}
