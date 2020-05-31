package walls_and_gates

const (
	inf  int = 2147483647 //represents a room
	door int = 0
	wall int = -1
)

// represents matrix coordinates
type node struct {
	row    int
	column int
}

// Store length of shortest path to a door in each inf, such that path does not go through other door or walls.
//
// Approach:
// - start by finding all doors
// - for each door, store distance of 1 in adjacent room (inf)
// - then start in rooms (inf) with distance 1; for each adjacent room with no distance stored, store distance of 2;
// - repeat until all rooms are filled
//
// Assumptions:
// - shortest distance from any room to any door is less than inf
// - array represents a rectangular matrix [m,n] with m>=0, n>=0
//
// Complexity (assuming array has m rows and n columns):
// - I think it is: O(5*m*n) = 1 pass over entire matrix + at most passing over each matrix element once and checking each neighbour.
func WallsAndGates(array [][]int) {
	workNodes := []node{}
	//find doors
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			if array[i][j] == door {
				workNodes = append(workNodes, node{row: i, column: j})
			}
		}
	}

	for len(workNodes) > 0 {
		nextWorkNodes := []node{}
		for _, w := range workNodes {
			for _, neighbour := range getNeighbours(w, len(array), len(array[0])) {
				if array[neighbour.row][neighbour.column] == inf {
					//hasn't been filled with distance yet, and isn't a wall or a door
					array[neighbour.row][neighbour.column] = array[w.row][w.column] + 1
					nextWorkNodes = append(nextWorkNodes, node{row: neighbour.row, column: neighbour.column})
				}
			}
		}
		workNodes = nextWorkNodes
	}
}

// return all neighbouring nodes for d for a m*n matrix
func getNeighbours(d node, m, n int) []node {
	result := []node{}
	if d.row+1 < m {
		result = append(result, node{row: d.row + 1, column: d.column})
	}
	if d.row-1 >= 0 {
		result = append(result, node{row: d.row - 1, column: d.column})
	}
	if d.column+1 < n {
		result = append(result, node{row: d.row, column: d.column + 1})
	}
	if d.column-1 >= 0 {
		result = append(result, node{row: d.row, column: d.column - 1})
	}
	return result
}
