package tracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"

	"github.com/shricodev/go-tracker-cli/cmd/utils"
)

// Add appends a new task to the list of trackers.
// It sets the task, marks it as incomplete, records the creation timestamp,
// and leaves the completion timestamp as NA.
func (t *Trackers) Add(task string) {
	tracker := Item{
		Task:        task,
		Completed:   false,
		CreatedAt:   time.Now().Format("Mon, 2006-01-02 15:04 PM"),
		CompletedAt: "---",
	}

	*t = append(*t, tracker)
}

// Complete marks a tracker as completed by its index and updates the completion timestamp.
// It returns an error if the index is invalid.
func (t *Trackers) Complete(index int) error {
	list := *t

	if index < 0 || index > len(list) {
		return &utils.InvalidIndexError{Index: index}
	}

	list[index-1].CompletedAt = time.Now().Format("Mon, 2006-01-02 15:04 PM")
	list[index-1].Completed = true
	return nil
}

// Delete removes a tracker from the list by its index.
// It returns an error if the index is invalid.
func (t *Trackers) Delete(index int) error {
	list := *t
	if index < 0 || index > len(list) {
		return errors.New("Invalid index")
	}

	*t = append(list[:index-1], list[index:]...)
	return nil
}

// LoadTrackers reads trackers from a JSON file and loads them into the Trackers slice.
// If the file does not exist, it creates a new one. It returns an error if the file
// has an invalid type or cannot be parsed.
func (t *Trackers) LoadTrackers(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		// Because if the file does not exist, we can simply create it
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("failed to load trackers: %w", err)
	}

	fileExtension := strings.ToLower(filepath.Ext(filename))

	if fileExtension != ".json" {
		return fmt.Errorf("Invalid file type: %s", fileExtension)
	}

	if len(file) == 0 {
		return errors.New("Provided file is empty")
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return fmt.Errorf("failed to unmarshall trackers: %w", err)
	}

	return nil
}

// Store writes the current state of the trackers to a JSON file.
// It returns an error if the marshalling process fails or if the file cannot be written.
func (t *Trackers) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("failed to marshall trackers: %w", err)
	}

	return os.WriteFile(filename, data, 0644)
}

// List prints a table representation of all trackers to the console, highlighting completed tasks.
// It includes headers for ID, Task, Completion status, Created At, and Completed At timestamps.
// The footer displays the total count of pending trackers.
func (t *Trackers) List() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "id"},
			{Align: simpletable.AlignCenter, Text: "Tracker"},
			{Align: simpletable.AlignCenter, Text: "Completed?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell
	for index, tracker := range *t {

		task := utils.Blue(tracker.Task)
		completed := utils.Red("no")
		createdAt := utils.Gray(tracker.CreatedAt)
		completedAt := utils.Gray(tracker.CompletedAt)

		if tracker.Completed {
			task = utils.Green(fmt.Sprintf("\u2705 %s", tracker.Task))
			completed = utils.Green("yes")
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", index+1)},
			{Text: task},
			{Text: completed},
			{Text: createdAt},
			{Text: completedAt},
		})
	}

	table.Body = &simpletable.Body{
		Cells: cells,
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: utils.Red(fmt.Sprintf("You have %d pending trackers", t.CountPending()))},
		},
	}

	table.SetStyle(simpletable.StyleRounded)
	table.Println()
}

// CountPending calculates and returns the total number of trackers that are not yet completed.
func (t *Trackers) CountPending() int {
	total := 0
	for _, tracker := range *t {
		if !tracker.Completed {
			total++
		}
	}
	return total
}
