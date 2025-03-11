package medium

import "fmt"

type TaskManager struct {
	tasks *TaskTree
	m     map[int]*Task // taskId - Task
}

type TaskTree struct {
	Val   *Task
	Left  *TaskTree
	Right *TaskTree
}

type Task struct {
	UserId   int
	TaskId   int
	Priority int
}

func (tr *TaskTree) Empty() bool {
	return tr == nil
}

func (tr *TaskTree) Add(task *Task) *TaskTree {
	if tr == nil {
		return &TaskTree{
			Val: task,
		}
	}

	node := tr
	for node != nil {
		if task.Priority > node.Val.Priority ||
			(task.Priority == node.Val.Priority && task.TaskId > node.Val.TaskId) {
			if node.Right != nil {
				node = node.Right
			} else {
				node.Right = &TaskTree{
					Val: task,
				}
				break
			}
		} else {
			if node.Left != nil {
				node = node.Left
			} else {
				node.Left = &TaskTree{
					Val: task,
				}
				break
			}
		}
	}

	return tr
}

func (tr *TaskTree) Edit(task *Task) *TaskTree {
	tr = tr.Rmv(task)
	return tr.Add(task)
}

func (tr *TaskTree) Rmv(task *Task) *TaskTree {
	if tr == nil {
		return nil
	}

	return tr
}

func (tr *TaskTree) Pop() *TaskTree {
	if tr == nil {
		return nil
	}

	node, right := tr, tr.Right
	for right != nil {
		node = right
		right = right.Right
	}

	node = tr
	for node.Right != nil {
		node = node.Right
	}

	//leftMost

	return node
}

// [userId, taskId, priority]
func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		tasks: nil,
		m:     map[int]*Task{},
	}

	for _, task := range tasks {
		tm.Add(task[0], task[1], task[2])
	}

	return tm
}

func print(tasks *TaskTree) {
	if tasks != nil {
		print(tasks.Right)
		fmt.Print("Priority: ", tasks.Val.Priority, ", TaskId: ", tasks.Val.TaskId, ", UserId: ", tasks.Val.UserId, "\n")
		print(tasks.Left)
	}
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	task := Task{
		UserId:   userId,
		TaskId:   taskId,
		Priority: priority,
	}
	tm.tasks = tm.tasks.Add(&task)
	tm.m[taskId] = &task
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {
	tm.tasks = tm.tasks.Edit(&Task{
		TaskId:   taskId,
		UserId:   tm.m[taskId].UserId,
		Priority: newPriority,
	})
	tm.m[taskId].Priority = newPriority
}

func (tm *TaskManager) Rmv(taskId int) {
	tm.tasks = tm.tasks.Rmv(tm.m[taskId])
}

func (tm *TaskManager) ExecTop() int {
	if tm.tasks.Empty() {
		return -1
	}

	pickedTask := tm.tasks.Pop()
	return pickedTask.Val.UserId
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
