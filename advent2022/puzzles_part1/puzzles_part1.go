package puzzles_part1

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

// combinations A,X(rock) = 1, B,Y(paper) = 2, C,Z(scissors) = 3, win = 6, draw = 3
var COMBO_FOR_CHOICE = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

// combinations A(rock) = 1, B(paper) = 2, C(scissors) = 3, win = 6, draw = 3, lose = 0, X = lose, Y = draw, Z = win
var COMBO_FOR_OUTCOME = map[string]int{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}

func rockPaperScissorsPuzzle(combos map[string]int) int {

	input := runPuzzleInput("./rockPaperScissors")
	total := 0
	for _, line := range input {
		total += combos[line]
	}
	return total
}

func getIntCharacterCode(c int32) int32 {
	value := int(c)
	if value >= 97 {
		return int32(value - 96)
	} else {
		return int32(value - 38)
	}
}

//func findBackpackBadges() int {
//	input := runPuzzleInput("./backpack")
//	var result int
//
//	for i := 0; i < len(input)/3; i++ {
//
//	}
//}

func findBackpackDuplicateItem() int {
	input := runPuzzleInput("./backpack")
	var result int

	for _, line := range input {
		firstHalf := map[int32]bool{}
		for _, c := range line[:len(line)/2] { //always even
			firstHalf[getIntCharacterCode(c)] = true
		}
		for _, c := range line[len(line)/2:] {
			rating := getIntCharacterCode(c)
			if firstHalf[rating] {
				result += int(rating)
				break
			}
		}
	}

	return result
}

func caloriesPuzzle() {
	input := runPuzzleInput("./puzzle_input")
	topMunchers := getElfWithTheMost(getCaloriesPerElf(input), 3)

	fmt.Printf("Top munchers %v\n", topMunchers)
	totalCalories := 0
	for _, munchMunch := range topMunchers {
		totalCalories += munchMunch.calories
	}
	fmt.Printf("Their total snacks is %d calories", totalCalories)
}

// returns the index of the elf with the most and total calories they have
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

func runPuzzleInput(path string) []string {
	result := []string{}

	file, err := os.Open(path)
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
