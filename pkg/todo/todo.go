package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"
)

type Task struct {
	Title       string
	CreatedAt   time.Time
	Completed   bool
	CompletedAt time.Time
}

type Tasks []Task

// Add appends a new Task to the list with the provided title, and returns it.
func (t *Tasks) Add(title string) (*Task, error) {
	var newT Task
	newT.Title = title
	newT.CreatedAt = time.Now()
	*t = append(*t, newT)
	return &newT, nil
}

// Complete marks the provided task title as done if found.
func (t *Tasks) Complete(title string) error {
	cT := t.Get(title)
	if *cT == (Task{}) {
		return errors.New("task not found")
	}
	cT.Completed = true
	cT.CompletedAt = time.Now()
	return nil
}

// Delete removes the provided task from the list if found.
func (t *Tasks) Delete(title string) error {

	tL := *t
	index := indexOfTitle(tL, title)
	if index == -1 {
		return errors.New("task not found")
	}
	dT := t.popTaskAtIndex(index)
	if dT == new(Task) {
		return errors.New("task could not be deleted")
	}
	return nil
}

// Get searches tasks by title and returns the corresponding task if found,
// or a new empty task if not.
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

// Save flushes the Tasks list into disk to the provided filename.
func (t *Tasks) Save(filename string) error {
	contents, err := json.Marshal(t)
	if err != nil {
		return err
	}
	perms := 0644 // owner: read, write. group read. others read.
	return os.WriteFile(filename, contents, fs.FileMode(perms))
}

// Load parses the Tasks list from the provided filename disk file.
func (t *Tasks) Load(filename string) error {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contents, t)
	if err != nil {
		return err
	}
	return nil
}

// String prints out a formatted task list. It implements the Stringer interface.
func (t *Tasks) String() string {
	formatted := ""
	for i, t := range *t {
		prefix := "[ ]"
		if t.Completed {
			prefix = "[X]"
		}
		formatted += fmt.Sprintf("%s %d: %s\n", prefix, i+1, t.Title)
	}
	return formatted
}

// indexOfTitle returns an index `int` value by searching current Tasks by Task.title.
// Returns `-1` when the provided title was not found in the data set.
func indexOfTitle(tL Tasks, title string) int {
	for i, task := range tL {
		if task.Title == title {
			return i
		}
	}
	return -1
}

// popTaskAtIndex returns task at provided index from the slice and returns it,
// if found, or a new empty one if not.
func (t *Tasks) popTaskAtIndex(i int) *Task {
	task := new(Task)
	tl := *t

	if len(tl) > i && i >= 0 {
		task = &tl[i]
	}

	*t = append((*t)[:i], (*t)[i+1:]...)

	return task
}
