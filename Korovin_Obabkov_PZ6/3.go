package main

import (
	"fmt"
	"sync"
)

type TaskQueue struct {
	tasks []string
	lock  sync.Mutex
}

// ////////////////////////
func (q *TaskQueue) Add(task string) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.tasks = append(q.tasks, task)
}

// ///////////////////////////
func (q *TaskQueue) Get() string {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.tasks) == 0 {
		return ""
	}

	task := q.tasks[0]

	q.tasks = q.tasks[1:]
	return task
}

func main() {
	queue := TaskQueue{}

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < 3; j++ {
				task := fmt.Sprintf("task-%d-%d", workerID, j)
				queue.Add(task)
				fmt.Printf("Worker %d added: %s\n", workerID, task)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Все задачи добавлены!")

	fmt.Println("Обработка задач:")
	for {
		task := queue.Get()
		if task == "" {
			break
		}
		fmt.Println("Обработано:", task)
	}

	fmt.Println("Все задачи обработаны!")
}
