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
	cT := t.Get(title)
	cT.Completed = true
	cT.CompletedAt = time.Now()
	return nil
}

func (t *Tasks) Delete(title string) error {

	tL := *t
	index := indexOfTitle(tL, title)
	if index == -1 {
		return errors.New("task index not found")
	}
	dT := t.popTaskAtIndex(index)
	if dT == new(Task) {
		return errors.New("task could not be deleted")
	}
	return nil
}

func (t *Tasks) Save(title string) error {
	return errors.New("not yet implemented")
}

// Get searches tasks by title and returns the task if found, or a new zeroed task if not.
func (t *Tasks) Get(title string) *Task {
	// Dereference the current pointer, as indexing is only supported on values.
	tL := *t

	g := new(Task)
	taskIndex := indexOfTitle(tL, title)

	if taskIndex > -1 {
		g = &tL[taskIndex]
	}

	return g
}

// IndexOfTitle returns an index `int` value by searching current Tasks by Task.title.
// Returns `-1` when the provided title was not found in the data set.
func indexOfTitle(tL Tasks, title string) int {
	for i, task := range tL {
		if task.Title == title {
			return i
		}
	}
	return -1
}

func (t *Tasks) popTaskAtIndex(i int) *Task {
	task := new(Task)
	tl := *t

	if len(tl) > i && i >= 0 {
		task = &tl[i]
	}

	*t = append((*t)[:i], (*t)[i+1:]...)

	return task
}
