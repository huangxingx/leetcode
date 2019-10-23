package minPathSum

func minPathSum(grid [][]int) int {

	for i := len(grid) - 1; i >= 0; i -- {
		for j := len(grid[0]) - 1; j >= 0; j -- {
			if i == len(grid)-1 && j != len(grid[0])-1 {
				grid[i][j] = grid[i][j] + grid[i][j+1]

			} else if i != len(grid)-1 && j == len(grid[0])-1 {
				grid[i][j] = grid[i][j] + grid[i+1][j]

			} else if i != len(grid)-1 && j != len(grid[0])-1 {
				grid[i][j] = grid[i][j] + min(grid[i+1][j], grid[i][j+1])
			}

		}
	}
	return grid[0][0]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
