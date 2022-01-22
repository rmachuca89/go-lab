package todo_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/rmachuca89/go-lab/pkg/todo"
)

func TestAdd(t *testing.T) {
	taskTitle := "Task 1"
	tL := new(todo.Tasks)

	t1, err := tL.Add(taskTitle)

	if err != nil {
		t.Fatalf("Unexpected error when calling Add() with %q: %q", taskTitle, err)
	}

	if t1.Title != taskTitle {
		t.Errorf("Expected %q, got %q instead.", taskTitle, t1.Title)
	}
}

func TestGet(t *testing.T) {
	want := &todo.Task{
		Title:     "Task 1",
		CreatedAt: time.Date(2022, time.January, 17, 9, 50, 0, 0, time.UTC),
	}
	tL := todo.Tasks{*want}
	taskTitle := "Task 1"

	got := tL.Get(taskTitle)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Tasks.Get(%q) got unexpected diff (-want +got):\n%s", taskTitle, diff)
	}

}

func TestComplete(t *testing.T) {
	t1 := todo.Task{
		Title:     "Task 1",
		CreatedAt: time.Date(2022, time.January, 17, 9, 50, 0, 0, time.UTC),
	}
	tL := todo.Tasks{t1}
	taskTitle := "Task 1"

	err := tL.Complete(taskTitle)
	if err != nil {
		t.Fatalf("Unexpected error when calling Complete() with %q: %q", taskTitle, err)
	}

	cT := tL.Get(taskTitle)

	if !cT.Completed {
		t.Errorf("Expected %q, got %q instead.", taskTitle, t1.Title)
	}
}

func TestDelete(t *testing.T) {
	tL := new(todo.Tasks)
	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _, t := range tasks {
		tL.Add(t)
	}

	taskTitle := "Task 2"
	if err := tL.Delete(taskTitle); err != nil {
		t.Fatalf("Unexpected error when calling Delete(%q): %q", taskTitle, err)
	}

	wantLen := 2
	if len(*tL) != wantLen {
		t.Errorf("Delete(%q) did not remove any task", taskTitle)
	}

	if dT := tL.Get(taskTitle); *dT != (todo.Task{}) {
		t.Errorf("Get(%q) stil found original task; returned: %v", taskTitle, dT)
	}
}

func TestSave(t *testing.T) {
	t.Skip("Not yet implemented.")
}

func TestLoad(t *testing.T) {
	t.Skip("Not yet implemented.")
}
