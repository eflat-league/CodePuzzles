package candy_crush

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCrush(t *testing.T) {
	tests := []struct {
		name      string
		input     [][]int
		exp       [][]int
		expStable bool
	}{
		{
			name: "corner",
			input: [][]int{
				{1, 1, 1},
				{1, 2, 2},
				{1, 2, 2},
			},
			exp: [][]int{
				{0, 0, 0},
				{0, 2, 2},
				{0, 2, 2},
			},
			expStable: false,
		},
		{
			name: "square",
			input: [][]int{
				{1, 1, 1},
				{1, 2, 1},
				{1, 1, 1},
			},
			exp: [][]int{
				{0, 0, 0},
				{0, 2, 0},
				{0, 0, 0},
			},
			expStable: false,
		},
		{
			name: "all",
			input: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			exp: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expStable: false,
		},
		{
			name: "horizontal",
			input: [][]int{
				{1, 1, 1},
				{1, 2, 1},
				{3, 1, 3},
				{1, 2, 1},
			},
			exp: [][]int{
				{0, 0, 0},
				{1, 2, 1},
				{3, 1, 3},
				{1, 2, 1},
			},
			expStable: false,
		},
		{
			name: "vertical",
			input: [][]int{
				{1, 2, 2},
				{1, 2, 1},
				{1, 1, 3},
				{2, 2, 1},
			},
			exp: [][]int{
				{0, 2, 2},
				{0, 2, 1},
				{0, 1, 3},
				{2, 2, 1},
			},
			expStable: false,
		},
		{
			name: "large",
			input: [][]int{
				{2, 2, 2, 1},
				{2, 2, 1, 1},
				{2, 1, 3, 1},
				{2, 2, 1, 3},
			},
			exp: [][]int{
				{0, 0, 0, 0},
				{0, 2, 1, 0},
				{0, 1, 3, 0},
				{0, 2, 1, 3},
			},
			expStable: false,
		},
		{
			name: "none match",
			input: [][]int{
				{1, 2, 2},
				{3, 2, 1},
				{1, 1, 3},
				{2, 2, 1},
			},
			exp: [][]int{
				{1, 2, 2},
				{3, 2, 1},
				{1, 1, 3},
				{2, 2, 1},
			},
			expStable: true,
		},
		{
			name: "none match with row of 0s",
			input: [][]int{
				{0, 0, 0},
				{1, 2, 2},
				{3, 2, 1},
				{1, 1, 3},
				{2, 2, 1},
			},
			exp: [][]int{
				{0, 0, 0},
				{1, 2, 2},
				{3, 2, 1},
				{1, 1, 3},
				{2, 2, 1},
			},
			expStable: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, stable := crush(test.input)
			if !reflect.DeepEqual(test.exp, got) {
				t.Fatalf("Test %s: \n(-want,+got)\n%s", test.name, cmp.Diff(test.exp, got))
			}
			if stable != test.expStable {
				t.Fatalf("Test %s: \n expected stable = %t\n got %t\n", test.name, test.expStable, stable)
			}
		})
	}
}

func TestDropCandy(t *testing.T) {
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
			name:  "slice of empty slices",
			input: [][]int{{}, {}},
			exp:   [][]int{{}, {}},
		},
		{
			name: "nothing to drop",
			input: [][]int{
				{1, 1, 2},
				{1, 1, 2},
				{2, 1, 2},
			},
			exp: [][]int{
				{1, 1, 2},
				{1, 1, 2},
				{2, 1, 2},
			},
		},
		{
			name: "already 0 on top",
			input: [][]int{
				{0, 0, 0},
				{0, 0, 2},
				{0, 0, 2},
			},
			exp: [][]int{
				{0, 0, 0},
				{0, 0, 2},
				{0, 0, 2},
			},
		},
		{
			name: "drop",
			input: [][]int{
				{1, 0, 0},
				{0, 2, 2},
				{0, 0, 2},
			},
			exp: [][]int{
				{0, 0, 0},
				{0, 0, 2},
				{1, 2, 2},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := dropCandy(test.input)
			if !reflect.DeepEqual(test.exp, got) {
				t.Fatalf("Test %s: \n(-want,+got)\n%s", test.name, cmp.Diff(test.exp, got))
			}
		})
	}
}

func TestStabilize(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		exp   [][]int
	}{
		{
			name: "most crushed",
			input: [][]int{
				{2, 1, 1},
				{2, 2, 2},
				{2, 1, 1},
				{2, 1, 1},
			},
			exp: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "leet code example",
			input: [][]int{
				{110, 5, 112, 113, 114},
				{210, 211, 5, 213, 214},
				{310, 311, 3, 313, 314},
				{410, 411, 412, 5, 414},
				{5, 1, 512, 3, 3},
				{610, 4, 1, 613, 614},
				{710, 1, 2, 713, 714},
				{810, 1, 2, 1, 1},
				{1, 1, 2, 2, 2},
				{4, 1, 4, 4, 1014},
			},
			exp: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{110, 0, 0, 0, 114},
				{210, 0, 0, 0, 214},
				{310, 0, 0, 113, 314},
				{410, 0, 0, 213, 414},
				{610, 211, 112, 313, 614},
				{710, 311, 412, 613, 714},
				{810, 411, 512, 713, 1014},
			},
		},
		{
			name:  "empty",
			input: [][]int{},
			exp:   [][]int{},
		},
		{
			name:  "empties", //can accept this as range is 3 < r < 50
			input: [][]int{{}, {}, {}},
			exp:   [][]int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Stabilize(test.input)
			if !reflect.DeepEqual(test.exp, got) {
				t.Fatalf("Test %s: \n(-want,+got)\n%s", test.name, cmp.Diff(test.exp, got))
			}
		})
	}
}
