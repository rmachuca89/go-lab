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
}

func tasksList(tL todo.Tasks) {
	if len(tL) == 0 {
		fmt.Println("There are no existing tasks! get to work...")
		return
	}
	fmt.Println("---------------------")
	fmt.Println("| No. | Title | Complete? |")
	fmt.Println("---------------------")
	for i, t := range tL {
		comp := "N"
		if t.Completed {
			comp = "Y"
		}
		fmt.Printf("%d, %q, %q\n", (i + 1), t.Title, comp)
	}
}

func taskAdd(title string, tL *todo.Tasks, dry bool) {
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

func main() {

	// 0. Init config and parse flags
	cfg := new(Config)
	flag.StringVar(&cfg.filename, "file", "tasks.json", "file containing the existing tasks in valid JSON format.")
	flag.BoolVar(&cfg.debug, "debug", false, "debug output")
	flag.BoolVar(&cfg.debug, "v", false, "verbose output (short)")
	flag.BoolVar(&cfg.debug, "verbose", false, "verbose output. alias to debug")
	flag.BoolVar(&cfg.dryRun, "dryrun", false, "perform a dry run of the task where no changes are performed)")
	flag.StringVar(&cfg.taskTitle, "title", "", "title for the task to operate on")

	flag.Parse()

	fmt.Println("Welcome to the TODO CLI APP!")

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
	// 1. Load all existing tasks.
	if err := tL.Load(cfg.filename); err != nil {
		log.Fatalf("Could not load tasks file: %q", err)
	}

	if len(os.Args) <= 1 {
		// 2. List tasks.
		tasksList(*tL)
	} else {
		// 3. Add tasks if args provided.
		if cfg.taskTitle == "" {
			log.Fatalln("the '--title' flag is required!")
		}
		if cfg.debug {
			log.Default().Printf("[DEBUG] Parsed new task title from args: %q", cfg.taskTitle)
		}
		taskAdd(cfg.taskTitle, tL, cfg.dryRun)

		// 4. Save new task to disk
		taskSave(tL, cfg.filename, cfg.dryRun)
		if cfg.debug {
			log.Default().Println("[DEBUG] new task saved to disk!")
		}
	}
}
