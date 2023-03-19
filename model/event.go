package model

import "fmt"

type Event struct {
	//Event ID
	EventID string

	//Full name of the Event
	EventName string

	//Internet address (URL) to a web site providing details
	ResourceLink string
}

func (event Event) String() string {
	return fmt.Sprintf("%s (%s)\nLink: %s", event.EventName, event.EventID, event.ResourceLink)
}
