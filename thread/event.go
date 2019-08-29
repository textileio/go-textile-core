package thread

import (
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-ipld-format"
)

type Event interface {
	format.Node

	Header() EventHeader
	Body() format.Node
}

type EventHeader interface {
	format.Node

	Schema() format.Node
	Doc() cid.Cid
}
