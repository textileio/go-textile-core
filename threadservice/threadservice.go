package threadservice

import (
	"context"

	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/go-textile-core/thread"
	tstore "github.com/textileio/go-textile-core/threadstore"
)

// Threadservice is an API for working with threads.
type Threadservice interface {
	// Threadstore persists thread log details.
	tstore.Threadstore

	// Host provides a network identity.
	Host() host.Host

	// DAGService provides a DAG API for reading and writing thread logs.
	DAGService() format.DAGService

	// Add data to a thread. Creates a new thread and own log if they don't exist.
	Add(ctx context.Context, body format.Node, opts ...AddOption) (peer.ID, cid.Cid, error)

	// Put an existing node to a log.
	Put(ctx context.Context, node thread.Node, opts ...PutOption) error

	// Pull paginates thread log events.
	Pull(ctx context.Context, t thread.ID, l peer.ID, opts ...PullOption) ([]thread.Node, error)

	// NewInvite returns a bundle of logs for the given thread.
	NewInvite(t thread.ID, reader bool) (format.Node, error)

	// Delete a thread.
	Delete(context.Context, thread.ID) error
}
