package infrastructure

import "github.com/textileio/go-textile-core/thread"

// EventStore provides an interface for storing events as sequences of indexed items.
// Event store appends events to stored sequences. It uses a record manager to map indexed items to database records,
// and it uses an item mapper to map indexed items to thread-level events.
type EventStore interface {
	ItemMapper
	RecordManager

	// Store appends thread event(s) to their sequence.
	PutEvents(event thread.Event) error
	// Returns all events for given sequence ID. Can be used to get events for a given aggregate.
	GetEvents(id string) (thread.Event, error)
}
