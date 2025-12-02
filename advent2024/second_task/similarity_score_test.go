package second_task

import (
	"advent2024/first_task"
	"fmt"
	"testing"
)

func TestGetSimilarityScore(t *testing.T) {
	firstList, secondList, err := first_task.ReadInputFile("../first_task/input")
	if err != nil {
		panic(err)
	}
	fmt.Println(getSimilarityScore(firstList, secondList))
}
