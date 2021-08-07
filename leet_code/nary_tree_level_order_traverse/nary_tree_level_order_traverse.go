package nary_tree_level_order_traverse

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	nextLevel := []*Node{root}
	for true {
		nextLevel = singleLevel(&result, nextLevel)
		if len(nextLevel) == 0 {
			break
		}
	}
	return result
}

//appends nodes of a single level to the result, in-place
//returns next level nodes
func singleLevel(result *[][]int, nodes []*Node) []*Node {
	levelNodes := []int{}
	nextLevelNodes := []*Node{}
	for _, node := range nodes {
		if node == nil {
			continue
		}
		nextLevelNodes = append(nextLevelNodes, node.Children...)
		levelNodes = append(levelNodes, node.Val)
	}
	if len(levelNodes) != 0 {
		*result = append(*result, levelNodes)
	}
	return nextLevelNodes
}
