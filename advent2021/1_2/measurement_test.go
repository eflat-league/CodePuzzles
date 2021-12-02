package measurement

import (
	"reflect"
	"testing"
)

func TestNumGreaterThanPrevious(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   int
	}{
		{
			name:  "empty",
			input: []int{},
			exp:   0,
		},
		{
			name:  "one",
			input: []int{1123},
			exp:   0,
		},
		{
			name:  "advent example",
			input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			exp:   7,
		},
	}

	for _, test := range tests {
		got := numGreaterThanPrevious(test.input)
		if got != test.exp {
			t.Fatalf("test %s: expected %d, got %d", test.name, test.exp, got)
		}
	}
}

func TestSlidingWindowOf3(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   []int
	}{
		{
			name:  "empty",
			input: []int{},
			exp:   []int{},
		},
		{
			name:  "one",
			input: []int{1123},
			exp:   []int{},
		},
		{
			name:  "two",
			input: []int{1123, 2},
			exp:   []int{},
		},
		{
			name:  "three",
			input: []int{3, 5, 9},
			exp:   []int{17},
		},
		{
			name:  "advent example",
			input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			exp:   []int{607, 618, 618, 617, 647, 716, 769, 792},
		},
	}

	for _, test := range tests {
		got := slidingWindowOf3(test.input)
		if !reflect.DeepEqual(test.exp, got) {
			t.Fatalf("test %s: expected %v, got %v", test.name, test.exp, got)
		}
	}
}

func TestRunPuzzleInput(t *testing.T) {
	runPuzzleInput()
}
