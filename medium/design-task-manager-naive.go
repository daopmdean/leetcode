package medium

type TaskManager struct {
	tasks []*Task
	// indexes [] // priority - []Task
}

type Task struct {
	UserId   int
	TaskId   int
	Priority int
}

func Constructor(tasks [][]int) TaskManager {
	ts := make([]*Task, 0, len(tasks))
	for _, task := range tasks {
		ts = append(ts, &Task{task[0], task[1], task[2]})
	}
	return TaskManager{
		tasks: ts,
	}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	this.tasks = append(this.tasks, &Task{
		UserId:   userId,
		TaskId:   taskId,
		Priority: priority,
	})
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	for _, tsk := range this.tasks {
		if tsk.TaskId == taskId {
			tsk.Priority = newPriority
		}
	}
}

func (this *TaskManager) Rmv(taskId int) {
	foundIdx := -1
	for i, task := range this.tasks {
		if task.TaskId == taskId {
			foundIdx = i
		}
	}

	if foundIdx == -1 {
		return
	}
	this.RmvByIdx(foundIdx)
}

func (this *TaskManager) RmvByIdx(taskIdx int) {
	this.tasks = append(this.tasks[:taskIdx], this.tasks[taskIdx+1:]...)
}

func (this *TaskManager) ExecTop() int {
	if len(this.tasks) == 0 {
		return -1
	}

	rmvIdx := 0
	pickedTaskId := this.tasks[0].TaskId
	max := this.tasks[0].Priority
	assignedUsers := map[int]int{
		pickedTaskId: this.tasks[0].UserId,
	} // taskId - userId
	for i, task := range this.tasks {
		if task.Priority < max {
			continue
		}

		if task.Priority > max {
			max = task.Priority
			pickedTaskId = task.TaskId
			rmvIdx = i
			assignedUsers[pickedTaskId] = task.UserId
			continue
		}

		if task.TaskId > pickedTaskId {
			pickedTaskId = task.TaskId
			rmvIdx = i
			assignedUsers[pickedTaskId] = task.UserId
		}
	}

	this.RmvByIdx(rmvIdx)
	return assignedUsers[pickedTaskId]
}

/**
* Your TaskManager object will be instantiated and called as such:
* obj := Constructor(tasks);
* obj.Add(userId,taskId,priority);
* obj.Edit(taskId,newPriority);
* obj.Rmv(taskId);
* param_4 := obj.ExecTop();
 */
