package median_two_sorted_arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		exp  float64
	}{
		//{
		//	name: "test1",
		//	arr1: []int{1, 3},
		//	arr2: []int{2},
		//	exp:  2.0,
		//},
		//{
		//	name: "test2",
		//	arr1: []int{1, 2},
		//	arr2: []int{3, 4},
		//	exp:  2.5,
		//},
		//{
		//	name: "test3",
		//	arr1: []int{1, 2},
		//	arr2: []int{3, 4, 5},
		//	exp:  3.0,
		//},
		//{
		//	name: "test4",
		//	arr1: []int{1, 2},
		//	arr2: []int{3, 4, 5, 6, 7, 8},
		//	exp:  4.5,
		//},
		//{
		//	name: "test5",
		//	arr1: []int{1},
		//	arr2: []int{},
		//	exp:  1,
		//},
		//{
		//	name: "test6",
		//	arr1: []int{},
		//	arr2: []int{1},
		//	exp:  1,
		//},
		{
			name: "test7",
			arr1: []int{1, 3},
			arr2: []int{2, 7},
			exp:  2.5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findMedianSortedArrays(test.arr1, test.arr2)
			assert.Equal(t, test.exp, got)
		})
	}
}
