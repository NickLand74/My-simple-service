package repo

import "fmt"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Repository struct {
	tasks  map[int]Task
	nextID int
}

func NewRepository() *Repository {
	return &Repository{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

// CreateTask - создание задачи
func (r *Repository) CreateTask(task Task) (int, error) {
	task.ID = r.nextID
	r.nextID++
	r.tasks[task.ID] = task
	return task.ID, nil
}

// GetTask - получение задачи по ID
func (r *Repository) GetTask(id int) (*Task, error) {
	task, exists := r.tasks[id]
	if !exists {
		return nil, fmt.Errorf("task not found")
	}
	return &task, nil
}

// GetAllTasks - получить все задачи
func (r *Repository) GetAllTasks() ([]Task, error) {
	tasks := make([]Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// UpdateTask - обновление задачи
func (r *Repository) UpdateTask(id int, updatedTask Task) error {
	if _, exists := r.tasks[id]; !exists {
		return fmt.Errorf("task not found")
	}
	updatedTask.ID = id
	r.tasks[id] = updatedTask
	return nil
}

// DeleteTask - удаление задачи
func (r *Repository) DeleteTask(id int) error {
	if _, exists := r.tasks[id]; !exists {
		return fmt.Errorf("task not found")
	}
	delete(r.tasks, id)
	return nil
}
