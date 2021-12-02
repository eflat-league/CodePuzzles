package measurement

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func slidingWindowOf3(input []int) []int {
	result := []int{}
	if len(input) < 3 {
		return result
	}
	for i := 2; i < len(input); i++ {
		result = append(result, input[i-2]+input[i-1]+input[i])
	}
	return result
}

func numGreaterThanPrevious(input []int) int {
	total := 0
	if len(input) <= 1 {
		return 0
	}
	for index := range input[1:] {
		if input[index+1] > input[index] { //offset current index by 1 as we are looping staring at index 1
			total++
		}
	}
	return total
}

func runPuzzleInput() {
	file, err := os.Open("./puzzle_input")
	if err != nil {
		panic(err)
	}
	input := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Errorf("Error converting %s to an int: %v", line, err))
		}
		input = append(input, number)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("RESULT: %d\n", numGreaterThanPrevious(slidingWindowOf3(input)))
}
