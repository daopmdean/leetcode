package medium

import "fmt"

type TaskManager struct {
	tasks []*Task
	m     map[int]*Task // taskId - userId
}

type Task struct {
	UserId   int
	TaskId   int
	Priority int
}

// [userId, taskId, priority]
func Constructor(tasks [][]int) TaskManager {
	ts := make([]*Task, 0, len(tasks))
	tm := TaskManager{
		tasks: ts,
		m:     map[int]*Task{},
	}

	for _, task := range tasks {
		tm.Add(task[0], task[1], task[2])
	}

	return tm
}

func print(tasks []*Task) {
	for _, task := range tasks {
		fmt.Println(task.Priority, task.TaskId, task.UserId)
	}
}

// 0 1 2 3 4 5 6 7 8
// 1 2 3 4 5 6 7 8 9
// use for tasks length more than 0
func (tm *TaskManager) findInsertIndex(taskId int, priority int) int {
	foundIdx := -1
	start, end := 0, len(tm.tasks)-1
	for start < end { // 4
		checkIdx := (start + end) / 2
		if tm.tasks[checkIdx].Priority < priority {
			start = checkIdx + 1
		} else if tm.tasks[checkIdx].Priority > priority {
			end = checkIdx - 1
		} else {
			foundIdx = checkIdx
			break
		}
	}

	if start == end {
		if tm.tasks[start].Priority == priority {
			foundIdx = start
		}
	}

	if foundIdx == -1 {
		if tm.tasks[start].Priority < priority {
			return start + 1
		}

		return start
	}

	left, right := foundIdx, foundIdx
	for left > 0 {
		if tm.tasks[left-1].Priority == priority {
			left--
			continue
		}
		break
	}

	for right < len(tm.tasks)-1 {
		if tm.tasks[right+1].Priority == priority {
			right++
			continue
		}
		break
	}

	for i := left; i <= right; i++ {
		if tm.tasks[i].TaskId < taskId {
			continue
		}
		return i
	}

	return right + 1
}

func (tm *TaskManager) findPriorityRange(priority int) (int, int) {
	idxFound := -1
	start, end := 0, len(tm.tasks)-1
	for start < end {
		checkIdx := (start + end) / 2
		if tm.tasks[checkIdx].Priority < priority {
			start = checkIdx + 1
		} else if tm.tasks[checkIdx].Priority > priority {
			end = checkIdx - 1
		} else {
			idxFound = checkIdx
			break
		}
	}

	if start == end && tm.tasks[start].Priority == priority {
		idxFound = start
	}

	left, right := idxFound, idxFound
	for left > 0 {
		if tm.tasks[left-1].Priority == priority {
			left--
			continue
		}
		break
	}

	for right < len(tm.tasks)-1 {
		if tm.tasks[right+1].Priority == priority {
			right++
			continue
		}
		break
	}

	return left, right
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	newTask := Task{
		UserId:   userId,
		TaskId:   taskId,
		Priority: priority,
	}

	tm.m[taskId] = &newTask

	if len(tm.tasks) == 0 {
		tm.tasks = append(tm.tasks, &newTask)
		return
	}

	insertIdx := tm.findInsertIndex(taskId, priority)

	if insertIdx == 0 {
		tm.tasks = append([]*Task{&newTask}, tm.tasks...)
	} else if insertIdx == len(tm.tasks) {
		tm.tasks = append(tm.tasks, &newTask)
	} else {
		newTasks := make([]*Task, 0, len(tm.tasks)+1)
		newTasks = append(newTasks, tm.tasks[:insertIdx]...)
		newTasks = append(newTasks, &newTask)
		newTasks = append(newTasks, tm.tasks[insertIdx:]...)
		tm.tasks = newTasks
	}
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {
	tm.Rmv(taskId)
	tm.Add(tm.m[taskId].UserId, taskId, newPriority)
}

func (tm *TaskManager) Rmv(taskId int) {
	priority := tm.m[taskId].Priority
	left, right := tm.findPriorityRange(priority)
	for i := left; i <= right; i++ {
		if tm.tasks[i].TaskId == taskId {
			tm.RmvByIdx(i)
			return
		}
	}
}

func (tm *TaskManager) RmvByIdx(idx int) {
	tm.tasks = append(tm.tasks[:idx], tm.tasks[idx+1:]...)
}

func (tm *TaskManager) ExecTop() int {
	print(tm.tasks)
	if len(tm.tasks) == 0 {
		return -1
	}
	task := tm.tasks[len(tm.tasks)-1]
	tm.tasks = tm.tasks[:len(tm.tasks)-1]
	return task.UserId
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
