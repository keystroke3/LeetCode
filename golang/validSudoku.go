package main

func isValidSudoku(board [][]byte) bool {
	columnSeen := [9][9]int{}
	boxSeen := [3][9]int{}
	for i, row := range board {
		rowSeen := [9]int{}
		if i == 3 || i == 6 {
			boxSeen = [3][9]int{}
		}
		for j, num := range row {
			if num == '.' {
				continue
			}
			intCol := num - 49
			if columnSeen[i][intCol] != 0 || rowSeen[intCol] != 0 {
				return false
			}
			columnSeen[i][intCol] = 1
			rowSeen[intCol] = 1

			switch j {
			case 0, 1, 2:
				if boxSeen[0][intCol] != 0 {
					return false
				}
				boxSeen[0][intCol] = 1
			case 3, 4, 5:
				if boxSeen[1][intCol] != 0 {
					return false
				}
				boxSeen[1][intCol] = 1
			default:
				if boxSeen[2][intCol] != 0 {
					return true
				}
				boxSeen[2][intCol] = 1

			}
		}

	}
	return true
}
