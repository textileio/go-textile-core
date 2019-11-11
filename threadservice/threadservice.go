package threadservice

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-textile-core/options"
	"github.com/textileio/go-textile-core/thread"
	tstore "github.com/textileio/go-textile-core/threadstore"
)

// Threadservice is an API for working with threads.
type Threadservice interface {
	io.Closer

	// Host provides a network identity.
	Host() host.Host

	// DAGService provides a DAG API to the network.
	format.DAGService

	// Store persists thread details.
	Store() tstore.Threadstore

	// AddThread from a multiaddress.
	AddThread(ctx context.Context, addr ma.Multiaddr, opts ...options.AddOption) (thread.Info, error)

	// PullThread for new records.
	PullThread(ctx context.Context, id thread.ID) error

	// Delete a thread.
	DeleteThread(ctx context.Context, id thread.ID) error

	// AddFollower to a thread.
	AddFollower(ctx context.Context, id thread.ID, pid peer.ID) error

	// AddRecord with body.
	AddRecord(ctx context.Context, id thread.ID, body format.Node) (Record, error)

	// GetRecord returns the record at cid.
	GetRecord(ctx context.Context, id thread.ID, rid cid.Cid) (thread.Record, error)

	// Subscribe returns a read-only channel of records.
	Subscribe(opts ...options.SubOption) Subscription
}

// Subscription receives thread record updates.
type Subscription interface {
	// Discard closes the subscription, disabling the reception of further records.
	Discard()

	// Channel returns the channel that receives records.
	Channel() <-chan Record
}

// Record wraps a thread.Record within a thread and log context.
type Record interface {
	// Value returns the underlying record.
	Value() thread.Record

	// ThreadID returns the record's thread ID.
	ThreadID() thread.ID

	// LogID returns the record's log ID.
	LogID() peer.ID
}
