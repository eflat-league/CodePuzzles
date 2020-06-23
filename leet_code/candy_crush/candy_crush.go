package candy_crush

func Stabilize(board [][]int) [][]int {
	result, stable := crush(board)
	for !stable {
		result = dropCandy(result)
		result, stable = crush(result)
	}
	return result
}

//given a rectangular 2d board, drops candy in each column down, treating 0 as empty space, and filling
//freed up space with 0
//does this in-place, modifying the board slice
func dropCandy(board [][]int) [][]int {
	if len(board) == 0 {
		return [][]int{}
	}
	result := zeroSlice(len(board), len(board[0]))
	//traverse each column  starting at the bottom and copy non-zero values to the result
	for col := 0; col < len(board[0]); col++ {
		resultRow := len(board) - 1
		for row := len(board) - 1; row >= 0; row-- {
			if board[row][col] != 0 {
				result[resultRow][col] = board[row][col]
				resultRow--
			}
		}
	}
	return result
}

//creates an m by n 2d slice initialized with 0
func zeroSlice(m, n int) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = 0
		}
	}
	return result
}

//creates a copy of a 2d slice
func copySlice(s [][]int) [][]int {
	result := make([][]int, len(s))
	for i, row := range s {
		result[i] = make([]int, len(s[0]))
		for j, value := range row {
			result[i][j] = value
		}
	}
	return result
}

//given a rectangular 2d board, crushes all candy that has 3 or more vertically or horizontally
//by replacing their value with 0
//returns crushed board and false if any candy was crushed and board is not stable, true if nothing was crushed and board is stable
func crush(board [][]int) ([][]int, bool) {
	if len(board) == 0 || len(board[0]) == 0 {
		return [][]int{}, true
	}

	//make a copy
	result := copySlice(board)
	boardIsStable := true

	//traverse vertically, crushing all candy >= 3 adjacent
	//we record crushed candy in the 'result' copy to handle 'corners' (eg 3 vertical connected to 3 horizontal)
	for col := 0; col < len(board[0]); col++ {
		counter := 1                  //how many adjacent same-color candy we saw
		currentColor := board[0][col] //current color being tracked
		for row := 1; row < len(board); row++ {
			if board[row][col] != currentColor {
				currentColor = board[row][col]
				counter = 1 //reset
				continue
			}

			counter += 1
			//if counter reached 3 or more, crush candy, but don't bother crushing 0s
			if currentColor != 0 {
				if counter == 3 {
					boardIsStable = false
					for crush := row; crush > row-counter; crush-- {
						result[crush][col] = 0 //crush this one and all before it
					}
				} else if counter > 3 {
					result[row][col] = 0 //crush this one
				}
			}
		}
	}

	//traverse horizontally
	//this is repetitive but can we avoid this without allocating more space?
	for row := 0; row < len(board); row++ {
		counter := 1                  //how many adjacent same-color candy we saw
		currentColor := board[row][0] //current color being tracked
		for col := 1; col < len(board[0]); col++ {
			if board[row][col] != currentColor {
				currentColor = board[row][col]
				counter = 1 //reset
				continue
			}

			counter += 1
			//if counter reached 3 or more, crush candy, but don't bother crushing 0
			if currentColor != 0 {
				if counter == 3 {
					boardIsStable = false
					for crush := col; crush > col-counter; crush-- {
						result[row][crush] = 0 //crush this one and all before it
					}
				} else if counter > 3 {
					result[row][col] = 0 //crush this one
				}
			}
		}
	}

	return result, boardIsStable
}
