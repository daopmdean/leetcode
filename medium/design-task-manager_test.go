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
		{7, 26, 23},
		{6, 1, 23},
	})
	//tm.Rmv(4)
	print(tm.tasks)
	//tm.Edit(4, 48)
	fmt.Println()
}

func print(tasks *TaskTree) {
	if tasks != nil {
		print(tasks.Right)
		fmt.Print("Priority: ", tasks.Val.Priority, ", TaskId: ", tasks.Val.TaskId, ", UserId: ", tasks.Val.UserId, "\n")
		print(tasks.Left)
	}
}
