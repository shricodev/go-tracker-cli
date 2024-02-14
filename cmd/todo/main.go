package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shricodev/go-tracker-cli"
)

const (
	trackerFile = ".tracker.json"
)

func main() {
	add := flag.Bool("add", false, "Add a new tracker")
	complete := flag.Int("complete", 0, "Mark a Tracker completed")
	delete := flag.Int("delete", 0, "Delete a Tracker")
	list := flag.Bool("list", false, "List all the trackers")

	flag.Parse()

	trackers := &tracker.Trackers{}
	if err := trackers.LoadTrackers(trackerFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := tracker.GetUserInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		trackers.Add(task)
		err = trackers.Store(trackerFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *complete > 0:
		err := trackers.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = trackers.Store(trackerFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *delete > 0:
		err := trackers.Delete(*delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = trackers.Store(trackerFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *list:
		trackers.List()

	default:
		fmt.Fprintln(os.Stdout, "invalid argument passed")
		os.Exit(0)
	}
}
