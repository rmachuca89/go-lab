package todo

import (
	"errors"
	"time"
)

type task struct {
	uuid        string
	title       string
	description string
	createdAt   time.Time
	completed   bool
	completedAt time.Time
}

type Tasks []task

func (t *Tasks) List() (*Tasks, error) {
	return new(Tasks), errors.New("not yet implemented")
}

func (t *Tasks) Add(title, desc string) (*task, error) {
	return new(task), errors.New("not yet implemented")
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

func (t *Tasks) Get(title string) error {
	return errors.New("not yet implemented")
}
