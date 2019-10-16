package dispatcher

// Event is a generic structure for dispatching events via the Dispatcher.
type Event interface {
	Body() []byte
	Time() []byte
	EntityID() string
	Type() string
}

// Reducer defines a generic structure that can be used by a Dispatcher.
type Reducer interface {
	Reduce(event Event) error
}
