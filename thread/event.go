package thread

import (
	"github.com/ipfs/go-ipld-format"
)

// Event is the Block format used by Textile threads
type Event interface {
	format.Node

	Header() EventHeader
	Body() format.Node

	Decrypt() (format.Node, error)
}

// EventHeader is the format of the event's header object
type EventHeader interface {
	format.Node

	Time() int
	Key() []byte
}
