package day1_puzzle1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	index    int
	calories int
}

func solvePuzzle() {
	input := runPuzzleInput()
	topMunchers := getElfWithTheMost(getCaloriesPerElf(input), 3)

	fmt.Printf("Top munchers %v\n", topMunchers)
	totalCalories := 0
	for _, munchMunch := range topMunchers {
		totalCalories += munchMunch.calories
	}
	fmt.Printf("Their total snacks is %d calories", totalCalories)
}

//returns the index of the elf with the most and total calories they have
func getElfWithTheMost(input []int, numElves int) []Elf {
	if len(input) == 0 {
		return nil
	}
	elves := []Elf{}
	for index, elfCalories := range input {
		elves = append(elves, Elf{
			index:    index,
			calories: elfCalories,
		})
	}
	//sort in reverse
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})
	if len(elves) < numElves {
		panic("Not enough elves")
	}
	return elves[0:numElves]
}

func getCaloriesPerElf(input []string) []int {
	result := []int{}
	newElf := true
	for _, line := range input {
		if line == "" {
			newElf = true
		} else {
			if newElf {
				result = append(result, 0)
				newElf = false
			}
			calories, _ := strconv.Atoi(line)
			result[len(result)-1] += calories
		}
	}
	return result
}

func runPuzzleInput() []string {
	result := []string{}

	file, err := os.Open("./puzzle_input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}
