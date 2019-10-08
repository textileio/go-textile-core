package threadservice

import (
	"context"

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
	Add(ctx context.Context, body format.Node, opts ...AddOption) (peer.ID, thread.Record, error)

	// Put an existing node to a log.
	Put(ctx context.Context, node thread.Record, opts ...PutOption) error

	// Pull paginates thread log events.
	Pull(ctx context.Context, t thread.ID, l peer.ID, opts ...PullOption) ([]thread.Record, error)

	// Logs returns info for each log in the given thread.
	Logs(t thread.ID) []thread.LogInfo

	// Delete a thread.
	Delete(context.Context, thread.ID) error
}
