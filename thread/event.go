package thread

import (
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-ipld-format"
)

// Event is the  Block format used by Textile threads
type Event interface {
	format.Node

	Header() EventHeader
	Body() format.Node
}

// EventHeader is the format of the event's header object
type EventHeader interface {
	format.Node

	Schema() format.Node
	Doc() cid.Cid
}
