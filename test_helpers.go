package events

import (
	"github.com/jmwri/go-events"
)

func eventIsPublishable(e go_events.Event) bool {
	_, ok := e.(go_events.PublishableEvent)
	return ok
}

func eventIsSubscribable(e go_events.Event) bool {
	_, ok := e.(go_events.SubscribableEvent)
	return ok
}
