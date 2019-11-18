package store

import (
	"github.com/google/uuid"
	ipldformat "github.com/ipfs/go-ipld-format"
	ds "github.com/ipfs/go-datastore"
)

const (
	// EmptyEntityID represents an empty EntityID
	EmptyEntityID = EntityID("")
)

// EntityID is the type used in models identities
type EntityID string

// NewEntityID generates a new identity for an instance
func NewEntityID() EntityID {
	return EntityID(uuid.New().String())
}

func (e EntityID) String() string {
	return string(e)
}

// Event is a local or remote event generated in a model and dispatcher
// by Dispatcher.
type Event interface {
	Body() []byte
	Time() []byte
	EntityID() EntityID
	Type() string
	Node() (ipldformat.Node, error)
}

// ActionType is the type used by actions done in a txn
type ActionType int

const (
	// Create indicates the creation of an instance in a txn
	Create ActionType = iota
	// Save indicates the mutation of an instance in a txn
	Save
	// Delete indicates the deletion of an instance by ID in a txn
	Delete
)

// Action is a operation done in the model
type Action struct {
	// Type of the action
	Type ActionType
	// EntityID of the instance in action
	EntityID EntityID
	// EntityType of the instance in action
	EntityType string
	// Previous is the instance before the action
	Previous interface{}
	// Current is the instance after the action was done
	Current interface{}
}

// EventCodec transforms actions generated in models to
// events dispatched to thread logs, and viceversa.
type EventCodec interface {
	// Reduce applies generated events into state
	Reduce(e Event, datastore ds.Datastore, baseKey ds.Key) error
	// Create corresponding events to be dispatched
	Create(ops []Action) ([]Event, error)
	EventFromBytes(data []byte) (Event, error)
}
