package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/rmachuca89/go-lab/pkg/todo"
)

const (
	errorTag  string = "[ERROR]"
	debugTag  string = "[DEBUG]"
	infoTag   string = "[INFO]"
	dryRunTag string = "[DRY-RUN]"
)

type Config struct {
	binName   string
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

	cfg.binName = path.Base(os.Args[0])
	fs.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for learning purposes.\n", cfg.binName)
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2022\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information:")
		fs.PrintDefaults()
	}
}

func tasksList(tL *todo.Tasks) {
	if len(*tL) == 0 {
		fmt.Println(infoTag, "There are no existing tasks! get to work...")
		return
	}
	fmt.Print(tL)
}

func taskAdd(tL *todo.Tasks, title string, dry bool) {
	if title == "" {
		log.Fatalln(errorTag, "New task title can not be empty.")
	}

	if dry {
		log.Default().Printf("%s Task with title %q would be attempted to be added.", debugTag, title)
		return
	}

	_, err := tL.Add(title)
	if err != nil {
		log.Fatalf("%s Could not add new task (%q): %q", errorTag, title, err)
	}
	log.Default().Println(infoTag, "Success. New task created.")
}

func taskSave(tL *todo.Tasks, filename string, dry bool) {
	if dry {
		log.Default().Println(dryRunTag, "Tasks would be attempted to be saved to disk.")
		return
	}

	if err := tL.Save(filename); err != nil {
		log.Fatalf("%s Could not save tasks to file: %q", errorTag, err)
	}
}

func taskComplete(tL *todo.Tasks, title string, dry bool) {
	if title == "" {
		log.Fatalln(errorTag, "Task title to complete required.")
	}

	if dry {
		log.Default().Println(dryRunTag, "Tasks would be attempted to be saved to disk.")
		return
	}

	if err := tL.Complete(title); err != nil {
		log.Fatalf("%s Could not mark task as complete: %q", errorTag, err)
	}
	log.Default().Println(infoTag, "Success. Task marked as complete.")
}

func main() {
	// 0. Init config and parse flags
	cfg := new(Config)
	fs := flag.NewFlagSet("todo flagset", flag.ExitOnError)
	cfg.RegisterFlags(fs)
	fs.Parse(os.Args[1:])

	if cfg.debug {
		log.Default().Printf("%s App Config: %+v", debugTag, cfg)
	}

	tL := new(todo.Tasks)
	// 1. Check if file exists; else create an empty
	if _, err := os.Stat(cfg.filename); err != nil {
		if cfg.debug {
			log.Default().Printf("%s File %q did not exist. Creating it...\n", debugTag, cfg.filename)
		}
		wErr := tL.Save(cfg.filename)
		if wErr != nil {
			log.Fatalf("%s Could not write initial empty file: %q", errorTag, err)
		}
	}
	if err := tL.Load(cfg.filename); err != nil {
		log.Fatalf("%s Could not load tasks file: %q", errorTag, err)
	}

	// Flags parsing
	switch {

	case cfg.complete:
		taskComplete(tL, cfg.taskTitle, cfg.dryRun)

	case cfg.taskTitle != "":
		taskAdd(tL, cfg.taskTitle, cfg.dryRun)

	default:
		tasksList(tL)
	}

	taskSave(tL, cfg.filename, cfg.dryRun)
	if cfg.debug {
		log.Default().Println(debugTag, "Tasks saved to disk.")
	}
}
