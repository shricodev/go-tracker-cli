## Go Tracker CLI

Go Tracker CLI is a command-line interface application designed to manage personal tasks or trackers.
With this tool, you can easily add, mark as completed, delete, and list your tasks directly from the terminal.

### Features

- **Add Trackers**: Quickly add new tasks to your list with a single command.
- **Complete Trackers**: Mark tasks as completed and track your progress efficiently.
- **Delete Trackers**: Remove tasks from your list when they are no longer needed.
- **List Trackers**: View all your active tasks in a neat table format.
- **Persistent Storage**: Your tasks are saved in a JSON file, ensuring they persist across sessions.
- **User-Friendly Interface**: Command-line interactions are straightforward and intuitive.

---

> 🍎 **NOTE**: As I have now completely moved into CLI. Most of my applications are in the CLI itself. This is just my way of replacing Microsoft Tracker. Just a small tracker CLI app, nothing **fancy**.

---

### Installation

To install Go Tracker CLI, clone the repository and build the binary using Go:

```bash
git clone git@github.com:shricodev/go-tracker-cli.git
cd go-tracker-cli
go build # OR, go install to install it globally on your system and
         # use it everywhere without being path limited. eg: go-tracker-cli --list
```

### Usage

```bash
./go-tracker-cli --add "Finish the report" # Add a tracker
./go-tracker-cli -complete=1               # Mark the tracker with idx 1 as completed
./go-tracker-cli --list                    # List all the trackers
./go-tracker-cli --delete=1                # Delete the tracker with idx 1
```
