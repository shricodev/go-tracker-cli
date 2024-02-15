package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shricodev/go-tracker-cli/cmd/trackers"
	"github.com/shricodev/go-tracker-cli/cmd/utils"
)

// Main handles the command-line interface for the tracker CLI. It accepts various
// flags to add, complete, delete, or list trackers. The trackers are loaded from
// and stored to a JSON file named '.tracker.json'.
func main() {
	const trackerFile = ".tracker.json"

	add := flag.Bool("add", false, "Add a new tracker")
	complete := flag.Int("complete", 0, "Mark a Tracker completed")
	delete := flag.Int("delete", 0, "Delete a Tracker")
	list := flag.Bool("list", false, "List all the trackers")

	flag.Parse()

	trackers := &tracker.Trackers{}
	utils.HandleGenericError(trackers.LoadTrackers(trackerFile))

	switch {
	case *add:
		task, err := utils.GetUserInput(os.Stdin, flag.Args()...)
		utils.HandleGenericError(err)
		trackers.Add(task)
		utils.HandleGenericError(trackers.Store(trackerFile))

	case *complete > 0:
		utils.HandleGenericError(trackers.Complete(*complete))
		utils.HandleGenericError(trackers.Store(trackerFile))

	case *delete > 0:
		utils.HandleGenericError(trackers.Delete(*delete))
		utils.HandleGenericError(trackers.Store(trackerFile))

	case *list:
		trackers.List()

	default:
		fmt.Fprintln(os.Stdout, "invalid argument passed")
		os.Exit(0)
	}
}
