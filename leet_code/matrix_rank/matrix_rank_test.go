package matrix_rank

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMatrixRankTransform(t *testing.T) {
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
			name:  "all ones",
			input: [][]int{{1, 1}, {1, 1}},
			exp:   [][]int{{1, 1}, {1, 1}},
		},
		{
			name:  "all sevens",
			input: [][]int{{7, 7}, {7, 7}},
			exp:   [][]int{{1, 1}, {1, 1}},
		},
		{
			name:  "negatives and equals",
			input: [][]int{{-7, -7}, {-7, 7}},
			exp:   [][]int{{1, 1}, {1, 2}},
		},
		{
			name:  "leet code example 3",
			input: [][]int{{20, -21, 14}, {-19, 4, 19}, {22, -47, 24}, {-19, 4, 19}},
			exp:   [][]int{{4, 2, 3}, {1, 3, 4}, {5, 1, 6}, {1, 3, 4}},
		},
		{
			name:  "leet code example 4",
			input: [][]int{{7, 3, 6}, {1, 4, 5}, {9, 8, 2}},
			exp:   [][]int{{5, 1, 4}, {1, 2, 3}, {6, 3, 1}},
		},
		{
			name:  "leet code fail 1",
			input: [][]int{{-37, -26, -47, -40, -13}, {22, -11, -44, 47, -6}, {-35, 8, -45, 34, -31}, {-16, 23, -6, -43, -20}, {47, 38, -27, -8, 43}},
			exp:   [][]int{{3, 4, 1, 2, 7}, {9, 5, 3, 10, 8}, {4, 6, 2, 7, 5}, {7, 9, 8, 1, 6}, {12, 10, 4, 5, 11}},
		},
		{
			name:  "leet code fail 2",
			input: [][]int{{-37, -50, -3, 44}, {-37, 46, 13, -32}, {47, -42, -3, -40}, {-17, -22, -39, 24}},
			exp:   [][]int{{2, 1, 4, 6}, {2, 6, 5, 4}, {5, 2, 4, 3}, {4, 3, 1, 5}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := matrixRankTransform(test.input)
			if !reflect.DeepEqual(test.exp, got) {
				t.Fatalf("%s: (-want,+got):\n%s\n", test.name, cmp.Diff(test.exp, got))
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		exp  int
	}{
		{
			name: "equal",
			a:    5,
			b:    5,
			exp:  5,
		},
		{
			name: "less",
			a:    -50,
			b:    5,
			exp:  5,
		},
		{
			name: "greater",
			a:    50,
			b:    5,
			exp:  50,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := max(test.a, test.b)
			if !reflect.DeepEqual(test.exp, got) {
				t.Fatalf("%s: (-want,+got):\n%s\n", test.name, cmp.Diff(test.exp, got))
			}
		})
	}
}

func TestRankSlice(t *testing.T) {
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
			input: []int{1234},
			exp:   []int{1},
		},
		{
			name:  "many",
			input: []int{4, 13, 1, -20, 50, -5},
			exp:   []int{4, 5, 3, 1, 6, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := rankSlice(test.input)
			if !reflect.DeepEqual(test.exp, got) {
				t.Fatalf("%s: (-want,+got):\n%s\n", test.name, cmp.Diff(test.exp, got))
			}
		})
	}
}

func TestDeduplicate(t *testing.T) {
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
			input: []int{1},
			exp:   []int{1},
		},
		{
			name:  "all same",
			input: []int{0, 0, 0, 0},
			exp:   []int{0},
		},
		{
			name:  "generic",
			input: []int{-20, -5, -5, 0, 0, 1, 1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 10101, 10101, 33333, 55555, 55555},
			exp:   []int{-20, -5, 0, 1, 2, 3, 4, 5, 6, 7, 10101, 33333, 55555},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := deduplicate(test.input)
			if !reflect.DeepEqual(test.exp, got) {
				t.Fatalf("%s: (-want,+got):\n%s\n", test.name, cmp.Diff(test.exp, got))
			}
		})
	}
}
