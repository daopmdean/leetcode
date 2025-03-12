package medium

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

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	task := Task{
		UserId:   userId,
		TaskId:   taskId,
		Priority: priority,
	}
	tm.tasks = add(tm.tasks, &task)
	tm.m[taskId] = &task
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {
	tm.tasks = rmv(tm.tasks, tm.m[taskId])
	tm.m[taskId].Priority = newPriority
	tm.tasks = add(tm.tasks, tm.m[taskId])

}

func (tm *TaskManager) Rmv(taskId int) {
	tm.tasks = rmv(tm.tasks, tm.m[taskId])
}

func (tm *TaskManager) ExecTop() int {
	if tm.tasks == nil {
		return -1
	}

	tasks, pickedTask := pop(tm.tasks)
	tm.tasks = tasks
	return pickedTask.UserId
}

func add(tr *TaskTree, task *Task) *TaskTree {
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

func pop(tr *TaskTree) (*TaskTree, *Task) {
	if tr == nil {
		return nil, nil
	}

	node, right := tr, tr.Right
	for right != nil {
		node = right
		right = right.Right
	}

	tr = rmv(tr, node.Val)
	return tr, node.Val
}

func rmv(tr *TaskTree, task *Task) *TaskTree {
	if tr == nil {
		return nil
	}

	if task.Priority > tr.Val.Priority {
		tr.Right = rmv(tr.Right, task)
	} else if task.Priority < tr.Val.Priority {
		tr.Left = rmv(tr.Left, task)
	} else {
		if task.TaskId > tr.Val.TaskId {
			tr.Right = rmv(tr.Right, task)
		} else if task.TaskId < tr.Val.TaskId {
			tr.Left = rmv(tr.Left, task)
		} else {
			if tr.Left == nil {
				return tr.Right
			} else if tr.Right == nil {
				return tr.Left
			}

			node := tr.Left
			for node.Right != nil {
				node = node.Right
			}

			tr.Val = node.Val
			tr.Left = rmv(tr.Left, node.Val)
		}
	}

	return tr
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
