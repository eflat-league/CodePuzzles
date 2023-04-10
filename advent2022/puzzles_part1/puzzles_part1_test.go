package puzzles_part1

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestGetCaloriesPerElf(t *testing.T) {
	//some nice edgy cases
	tests := []struct {
		name  string
		input []string
		exp   []int
	}{
		{
			name:  "last elf has 0 calories",
			input: []string{"1000", "2000", "", "0"},
			exp:   []int{3000, 0},
		},
		{
			name:  "empty",
			input: []string{},
			exp:   []int{},
		},
		{
			name:  "empty lines",
			input: []string{"", "", ""},
			exp:   []int{},
		},
		{
			name:  "empty line start",
			input: []string{"", "2000", "1000"},
			exp:   []int{3000},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getCaloriesPerElf(test.input)
			if !reflect.DeepEqual(test.exp, got) {
				t.Fatalf("(+want,-got) %s", cmp.Diff(test.exp, got))
			}
		})
	}
}

func TestCaloriesPuzzle(t *testing.T) {
	caloriesPuzzle()
}

func TestRockPaperScissorsPuzzle(t *testing.T) {
	fmt.Println(rockPaperScissorsPuzzle(COMBO_FOR_CHOICE))
	fmt.Println(rockPaperScissorsPuzzle(COMBO_FOR_OUTCOME))
}

func TestGetIntCharacterCode(t *testing.T) {
	fmt.Println(getIntCharacterCode('z'))
	fmt.Println(getIntCharacterCode('Z'))
}

func TestFindBackpackDuplicateItem(t *testing.T) {
	fmt.Println(findBackpackDuplicateItem())
}
