package thread

import (
	"time"

	"github.com/ipfs/go-ipld-format"
	ic "github.com/libp2p/go-libp2p-core/crypto"
)

type Event interface {
	format.Node

	Header() EventHeader
	Body() format.Node
}

type EventHeader interface {
	format.Node

	Type() []byte
	Actor() ic.PubKey
	Doc() string
	Date() time.Time
}
