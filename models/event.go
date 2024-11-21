package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{
	{
		ID:          122,
		Name:        "Go Meet",
		Description: "just a simple desc",
		Location:    "KJoiolkat",
		DateTime:    time.Now(),
		UserID:      1,
	},
}

func main() {
	// Do something with events
}

func (e Event) Save() {
	events = append(events, e)
}

func GetEvents() []Event {
	return events
}
