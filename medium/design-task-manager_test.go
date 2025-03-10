package medium

import (
	"fmt"
	"testing"
)

func TestTM(t *testing.T) {
	tm := Constructor([][]int{
		{2, 12, 32},
		{3, 27, 33},
		{10, 5, 23},
		{8, 4, 3},
		//{2, 26, 15},
		//{6, 1, 19},
	})
	tm.Rmv(4)
	print(tm.tasks)
	//tm.Edit(4, 48)
	fmt.Println()
}
