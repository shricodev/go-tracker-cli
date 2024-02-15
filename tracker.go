package tracker

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Completed   bool
	CreatedAt   string
	CompletedAt string
}

type Trackers []item

func (t *Trackers) Add(task string) {
	tracker := item{
		Task:        task,
		Completed:   false,
		CreatedAt:   time.Now().Format(time.RFC822),
		CompletedAt: "---",
	}

	*t = append(*t, tracker)
}

func (t *Trackers) Complete(index int) error {
	list := *t

	if index < 0 || index > len(list) {
		return errors.New("Invalid index")
	}

	list[index-1].CompletedAt = time.Now().Format(time.RFC822)
	list[index-1].Completed = true
	return nil
}

func (t *Trackers) Delete(index int) error {
	list := *t
	if index < 0 || index > len(list) {
		return errors.New("Invalid index")
	}

	*t = append(list[:index-1], list[index:]...)
	return nil
}

func (t *Trackers) LoadTrackers(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		// Because if the file does not exist, we can simply create it
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
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
		return err
	}

	return nil
}

func (t *Trackers) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

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

		task := blue(tracker.Task)
		completed := red("no")
		createdAt := gray(tracker.CreatedAt)
		completedAt := gray(tracker.CompletedAt)

		if tracker.Completed {
			task = green(fmt.Sprintf("\u2705 %s", tracker.Task))
			completed = green("yes")
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
			{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending trackers", t.CountPendingTrackers()))},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func GetUserInput(reader io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	// When the user sends the data from a pipe command
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := strings.TrimSpace(scanner.Text())

	if len(text) == 0 {
		return "", errors.New("Cannot add an empty Tracker")
	}

	return text, nil
}

func (t *Trackers) CountPendingTrackers() int {
	total := 0
	for _, tracker := range *t {
		if !tracker.Completed {
			total++
		}
	}
	return total
}
