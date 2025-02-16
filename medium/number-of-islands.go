package medium

func numIslands(grid [][]byte) int {

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				markZero(grid, i, j)
				count++
			}
		}
	}

	return count
}

func markZero(grid [][]byte, i, j int) {
	if i < 0 || j < 0 ||
		i >= len(grid) || j >= len(grid[0]) ||
		grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	markZero(grid, i+1, j)
	markZero(grid, i-1, j)
	markZero(grid, i, j+1)
	markZero(grid, i, j-1)
}
