package main

import "fmt"

func main() {
	manager := TaskManager{}
	manager.AddTask("Сделать ДЗ")
	manager.AddTask("Написать конспект по GO")
	manager.CompleteTask("Сделать ДЗ")
	fmt.Println(manager.GetCompleted())
}

type Task struct {
	Title     string
	Completed bool
}

type TaskManager struct {
	Tasks []Task
}

func (tm *TaskManager) AddTask(title string) {
	tm.Tasks = append(tm.Tasks, Task{Title: title})
}

func (tm *TaskManager) CompleteTask(title string) {
	for i := range tm.Tasks {
		if tm.Tasks[i].Title == title {
			tm.Tasks[i].Completed = true
		}
	}
}

func (tm *TaskManager) RemoveTask(title string) {
	newTasks := []Task{}
	for _, t := range tm.Tasks {
		if t.Title != title {
			newTasks = append(newTasks, t)
		}
	}
	tm.Tasks = newTasks
}

func (tm *TaskManager) GetCompleted() []Task {
	var res []Task
	for _, t := range tm.Tasks {
		if t.Completed {
			res = append(res, t)
		}
	}
	return res
}
