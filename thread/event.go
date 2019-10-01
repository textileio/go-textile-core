package thread

import (
	"context"
	"time"

	"github.com/ipfs/go-ipld-format"
	"github.com/textileio/go-textile-core/crypto"
)

// Event is the Block format used by Textile threads
type Event interface {
	format.Node

	Header(context.Context, format.DAGService, crypto.DecryptionKey) (EventHeader, error)
	Body(context.Context, format.DAGService, crypto.DecryptionKey) (format.Node, error)
}

// EventHeader is the format of the event's header object
type EventHeader interface {
	format.Node

	Time() time.Time
	Key() (crypto.DecryptionKey, error)
}
