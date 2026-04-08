package structures

import (
	"errors"
	"sync"
	"time"
)

type TodoList struct {
	Tasks  []Task
	mtx    sync.Mutex
	nextId int
}

type Task struct {
	Id      int    `json:"id"`
	Author  string `json:"author"`
	Text    string `json:"text"`
	Created time.Time
	Ended   *time.Time
}

func (t *TodoList) CreateTask(author string, text string) ([]Task, error) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	t.nextId++
	if text == "" {
		return nil, errors.New("MSG is empty")
	}
	newTask := Task{
		Id:      t.nextId,
		Author:  author,
		Text:    text,
		Created: time.Now(),
	}

	t.Tasks = append(t.Tasks, newTask)

	return t.Tasks, nil
}

func (t *TodoList) ListTasks() []Task {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	return t.Tasks
}

func (t *TodoList) GetTaskById(id int) (Task, error) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	for _, task := range t.Tasks {
		if id == task.Id {
			return task, nil
		}
	}
	return Task{}, errors.New("Task not found")
}
