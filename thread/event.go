package thread

import (
	"time"

	"github.com/ipfs/go-ipld-format"
	"github.com/libp2p/go-libp2p-core/peer"
)

type Event interface {
	format.Node

	Header() EventHeader
	Body() format.Node
}

type EventHeader interface {
	format.Node

	Type() []byte
	Actor() peer.ID
	Doc() string
	Date() time.Time
}
