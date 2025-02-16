package medium

func numIslandsNotModifyOriginSlice(grid [][]byte) int {
	walked := map[int]map[int]bool{}
	for i := 0; i < len(grid); i++ {
		walked[i] = make(map[int]bool)
	}

	islands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !walked[i][j] {
				walkThroughIsland(grid, walked, i, j)
				islands++
			}
		}
	}

	return islands
}

func walkThroughIsland(grid [][]byte, walked map[int]map[int]bool, i, j int) {
	if i < 0 || i >= len(grid) ||
		j < 0 || j >= len(grid[0]) {
		return
	}

	if walked[i][j] {
		return
	}

	if grid[i][j] == '1' {
		walked[i][j] = true
	} else {
		return
	}

	walkThroughIsland(grid, walked, i+1, j)
	walkThroughIsland(grid, walked, i-1, j)
	walkThroughIsland(grid, walked, i, j+1)
	walkThroughIsland(grid, walked, i, j-1)
}
