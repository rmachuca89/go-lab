package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName                   = "todo"
	buildOutDir, buildOutFile string
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	buildDir, err := os.MkdirTemp(os.TempDir(), "go-tmp-build")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create temp build dir %s: %s", buildDir, err)
		os.Exit(2)
	}

	buildOutDir = buildDir
	buildOutFile = filepath.Join(buildDir, binName)
	fmt.Println("Build File: ", buildOutFile)
	build := exec.Command("go", "build", "-o", buildOutFile)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(2)
	}

	fmt.Println("Running tests...")
	exitStatus := m.Run()

	fmt.Println("Cleaning up...")
	if err := os.Remove(buildOutFile); err != nil {
		fmt.Fprintf(os.Stderr, "Could not delete tool binary %s: %s", binName, err)
		os.Exit(2)
	}
	if err := os.RemoveAll(buildDir); err != nil {
		fmt.Fprintf(os.Stderr, "Could not delete temporary dir %s: %s", buildDir, err)
		os.Exit(2)
	}
	os.Exit(exitStatus)
}

func TestTodoCLI(t *testing.T) {

	tt := "test task number 1"

	if buildOutFile == "" || buildOutDir == "" {
		t.Fatalf("no built binary (%s) available", binName)
	}
	// Change to our temporal app testing dir.
	os.Chdir(buildOutDir)

	t.Run("AddNewTask", func(t *testing.T) {
		args := []string{"--title", tt}
		cmd := exec.Command(buildOutFile, args...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		args := []string{"--title", tt, "--complete"}
		cmd := exec.Command(buildOutFile, args...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(buildOutFile)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		// Expected format provided by Tasks.String()
		want := "[X] 1: " + tt + "\n"

		got := string(out)
		if want != got {
			t.Errorf("Expected %q, got %q instead\n", want, got)
		}
	})
}
