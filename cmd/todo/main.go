package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rmachuca89/go-lab/pkg/todo"
)

type Config struct {
	filename  string
	debug     bool
	taskTitle string
	dryRun    bool
	complete  bool
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&cfg.filename, "file", "tasks.json", "file containing the existing tasks in valid JSON format.")
	fs.BoolVar(&cfg.debug, "debug", false, "debug output")
	fs.BoolVar(&cfg.debug, "v", false, "verbose output (short)")
	fs.BoolVar(&cfg.debug, "verbose", false, "verbose output. alias to debug")
	fs.BoolVar(&cfg.dryRun, "dryrun", false, "perform a dry run of the task where no changes are performed")
	fs.StringVar(&cfg.taskTitle, "title", "", "title for the task to operate on")
	fs.BoolVar(&cfg.complete, "complete", false, "completes the provided task title")
}

func tasksList(tL todo.Tasks) {
	if len(tL) == 0 {
		fmt.Println("There are no existing tasks! get to work...")
		return
	}
	for _, t := range tL {
		fmt.Println(t.Title)
	}
}

func taskAdd(tL *todo.Tasks, title string, dry bool) {
	if title == "" {
		log.Fatalln("New task title can not be empty!")
	}

	if dry {
		log.Default().Printf("[DRY RUN] Task with title %q would be attempted to be added", title)
		return
	}

	t, err := tL.Add(title)
	if err != nil {
		log.Fatalf("Could not add new task (%q): %q", title, err)
	}
	log.Default().Printf("New task (%q) added successfully", t.Title)
}

func taskSave(tL *todo.Tasks, filename string, dry bool) {
	if dry {
		log.Default().Println("[DRY RUN] Tasks would be attempted to be saved to disk")
		return
	}

	if err := tL.Save(filename); err != nil {
		log.Fatalf("Could not save tasks to file: %q", err)
	}
}

func taskComplete(tL *todo.Tasks, title string, dry bool) {
	if dry {
		log.Default().Println("[DRY RUN] Tasks would be attempted to be saved to disk")
		return
	}

	if err := tL.Complete(title); err != nil {
		log.Fatalf("Could not mark task as complete: %q", err)
	}
	log.Default().Printf("Task (%q) marked as complete.", title)
}

func main() {
	// 0. Init config and parse flags
	cfg := new(Config)
	fs := flag.NewFlagSet("todo flagset", flag.ExitOnError)
	cfg.RegisterFlags(fs)
	fs.Parse(os.Args[1:])

	if cfg.debug {
		log.Default().Printf("[DEBUG] App Config: %+v", cfg)
	}

	tL := new(todo.Tasks)
	// 1. Check if file exists; else create an empty
	if _, err := os.Stat(cfg.filename); err != nil {
		if cfg.debug {
			log.Default().Printf("[DEBUG] File %q did NOT exist\n", cfg.filename)
		}
		wErr := tL.Save(cfg.filename)
		if wErr != nil {
			log.Fatalf("Could not write initial empty file: %q", err)
		}
	}
	if err := tL.Load(cfg.filename); err != nil {
		log.Fatalf("Could not load tasks file: %q", err)
	}

	// Flags parsing
	switch {

	case cfg.taskTitle != "" && cfg.complete:
		taskComplete(tL, cfg.taskTitle, cfg.dryRun)

	case cfg.taskTitle != "":
		taskAdd(tL, cfg.taskTitle, cfg.dryRun)

	default:
		tasksList(*tL)
	}

	taskSave(tL, cfg.filename, cfg.dryRun)
	if cfg.debug {
		log.Default().Println("[DEBUG] new task saved to disk!")
	}
}
