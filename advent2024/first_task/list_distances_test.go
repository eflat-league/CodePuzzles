package first_task

import (
	"fmt"
	"testing"
)

func TestDistances(t *testing.T) {
	firstList, secondList, err := ReadInputFile("input")
	if err != nil {
		panic(err)
	}
	fmt.Println(GetDistance(firstList, secondList))
}
