package main
import (
	"fmt"
	"sync"
	
)
type Task struct {
	Id int
	Data string
}

type Queue struct {
	mu sync.Mutex
	tasks []Task
	nextID int
}

func NewQueue() *Queue {
	return &Queue{
		tasks: make([]Task, 0),
		nextID: 1,
	}
}

func (q*Queue) Add(data Task){
	q.mu.Lock()
	defer q.mu.Unlock()
	task := Task{
		ID: q.nextID,
		Data: data,
	}
	q.nextID++
	q.tasks=append(q.tasks,task)
	fmt.Printf("Добавлено задание: %d - %s\n", task.ID,task.Data)
}

func (q *Queue) Get() (Task, bool) {
	q.mu.Lockq()
	defer q.mu.Unlock()
	if len(q.tasks) == 0{
		return Task{}, false
	}
	task := q.tasks[0]
	q.task = q.tasks[1:]
	fmt.Printf("обрабатвается задача: %d - %s \n",task.ID, task.Data)
	return task, true
}

func main(){
	queue := NewQueue()

	for i:= 0; i<3; i++{
		go func(workerID int) {
			for {
				task, ok := queue.Get()
				if ok {
					time.Sleep(200 * time.Millisecond)
				} else {
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(i)
 	}
	for i := 0; i <10; i++ {
		go func (num int) {
			data := fmt.Sprintf("task_data_%d", num)
			queue.Add(data)
			time.Sleep(50 * time.Millisecond)
		}(i)
	}
	time.Sleep(2*time.Second)
}