package nary_tree_level_order_traverse

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		exp  [][]int
	}{
		{
			name: "nil",
			root: nil,
			exp:  [][]int{},
		},
		{
			name: "one node",
			root: &Node{
				Val: 12,
			},
			exp: [][]int{{12}},
		},
		{
			name: "four levels with nil",
			root: &Node{
				Val: 12,
				Children: []*Node{
					{
						Val: 1,
						Children: []*Node{
							nil,
							{Val: 3},
						},
					},
					{
						Val: 2,
						Children: []*Node{
							{Val: 4},
							{
								Val: 5,
								Children: []*Node{
									{Val: 6},
								},
							},
						},
					},
				},
			},
			exp: [][]int{{12}, {1, 2}, {3, 4, 5}, {6}},
		},
		{
			name: "node with all nil children",
			root: &Node{
				Val:      12,
				Children: []*Node{nil, nil},
			},
			exp: [][]int{{12}}, //verifies we do not append empty slice
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := levelOrder(test.root)
			if !reflect.DeepEqual(test.exp, got) {
				t.Fatalf("%s : (-want, +got) %s", test.name, cmp.Diff(test.exp, got))
			}
		})
	}
}
