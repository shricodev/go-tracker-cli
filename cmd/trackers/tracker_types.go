package tracker

type Item struct {
	Task        string
	Completed   bool
	CreatedAt   string
	CompletedAt string
}

type Trackers []Item
