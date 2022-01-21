package todo

import (
	"errors"
	"time"
)

type Task struct {
	Title       string
	CreatedAt   time.Time
	Completed   bool
	CompletedAt time.Time
}

type Tasks []Task

func (t *Tasks) Add(title string) (*Task, error) {
	var newT Task
	newT.Title = title
	newT.CreatedAt = time.Now()
	*t = append(*t, newT)
	return &newT, nil
}

func (t *Tasks) Complete(title string) error {
	return errors.New("not yet implemented")
}

func (t *Tasks) Delete(title string) error {
	return errors.New("not yet implemented")
}

func (t *Tasks) Save(title string) error {
	return errors.New("not yet implemented")
}

func (t *Tasks) Get(title string) *Task {
	return new(Task)
}
