package walls_and_gates

import (
	"reflect"
	"testing"
)

func TestWallsAndGates(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		exp   [][]int
	}{
		{
			name:  "empty",
			input: [][]int{},
			exp:   [][]int{},
		},
		{
			name:  "one door",
			input: [][]int{{door}},
			exp:   [][]int{{door}},
		},
		{
			name: "one door",
			input: [][]int{
				{wall, door, wall},
				{wall, wall, wall},
				{inf, inf, inf},
			},
			exp: [][]int{
				{wall, door, wall},
				{wall, wall, wall},
				{inf, inf, inf},
			},
		},
		{
			name: "leet code example",
			input: [][]int{
				{inf, wall, door, inf},
				{inf, inf, inf, wall},
				{inf, wall, inf, wall},
				{door, wall, inf, inf},
			},
			exp: [][]int{
				{3, wall, door, 1},
				{2, 2, 1, wall},
				{1, wall, 2, wall},
				{door, wall, 3, 4},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			WallsAndGates(test.input)
			if !reflect.DeepEqual(test.exp, test.input) {
				t.Fatalf("Test %s:\n(want)\n%v\n(got):\n%v\n", test.name, test.exp, test.input)
			}
		})
	}
}
