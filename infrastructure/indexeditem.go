package infrastructure

// IndexedItem represents a generic 'item' that has a specific position in an event log.
type IndexedItem interface {
	// Identity of the sequence to which this item belongs.
	ID() string
	// Position of the item in its sequence. This could be a hybrid logical clock value.
	// e.g. 1568419140
	Position() int64
	// The 'topic' to which this item pertains. This could be a thread id or domain even type.
	// e.g. (V1, 256 bit) textile.thread.bafyoiobghzefwlidfrwkqmzz2ka66zgmdmgeobw2mimktr5jivsavya
	Topic() string
	// The actual state of the indexed item. Currently not clear what this needs to be/look like?
	State() interface{}
}
