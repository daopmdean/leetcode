package medium

import "fmt"

type TaskManager struct {
	tasks *TaskLink
	m     map[int]*Task // taskId - Task
}

type TaskLink struct {
	Val  *Task
	Next *TaskLink
}

type Task struct {
	UserId   int
	TaskId   int
	Priority int
}

func (tl *TaskLink) Empty() bool {
	return tl == nil
}

func (tl *TaskLink) Add(task *Task) *TaskLink {
	if tl == nil {
		return &TaskLink{
			Val: task,
		}
	}

	if task.Priority > tl.Val.Priority ||
		(task.Priority == tl.Val.Priority && task.TaskId > tl.Val.TaskId) {
		return &TaskLink{
			Val:  task,
			Next: tl,
		}
	}

	first, sec := tl, tl.Next
	for sec != nil {
		if task.Priority > sec.Val.Priority ||
			(task.Priority == sec.Val.Priority && task.TaskId > sec.Val.TaskId) {
			first.Next = &TaskLink{
				Val:  task,
				Next: sec,
			}
			return tl
		} else {
			first = sec
			sec = sec.Next
		}
	}

	first.Next = &TaskLink{
		Val: task,
	}

	return tl
}

func (tl *TaskLink) Edit(task *Task) *TaskLink {
	tl = tl.Rmv(task.TaskId)
	return tl.Add(task)
}

func (tl *TaskLink) Rmv(taskId int) *TaskLink {
	if tl == nil {
		return nil
	}

	first, sec := tl, tl.Next
	if first.Val.TaskId == taskId {
		return sec
	}

	for sec != nil {
		if sec.Val.TaskId == taskId {
			first.Next = sec.Next
			return tl
		}
		first = sec
		sec = sec.Next
	}

	return tl
}

func (tl *TaskLink) Pop() *TaskLink {
	if tl == nil {
		return nil
	}
	return tl.Next
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

func print(tasks *TaskLink) {
	for tasks != nil {
		fmt.Print("Priority: ", tasks.Val.Priority, ", TaskId: ", tasks.Val.TaskId, ", UserId: ", tasks.Val.UserId, "\n")
		tasks = tasks.Next
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
}

func (tm *TaskManager) Rmv(taskId int) {
	tm.tasks = tm.tasks.Rmv(taskId)
}

func (tm *TaskManager) ExecTop() int {
	if tm.tasks.Empty() {
		return -1
	}

	pickedTask := tm.tasks.Val
	tm.tasks = tm.tasks.Pop()
	return pickedTask.UserId
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
