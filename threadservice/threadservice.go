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

	// Host provides a network listener identity.
	Host() host.Host

	// DAGService provides a DAG API for reading and writing thread logs.
	DAGService() format.DAGService

	// Add a new record by wrapping body. See AddOption for more.
	Add(ctx context.Context, body format.Node, opts ...AddOption) (*Record, error)

	// Put an existing record. See PutOption for more.
	Put(ctx context.Context, node thread.Record, opts ...PutOption) error

	// Get returns the record at cid.
	Get(ctx context.Context, id thread.ID, lid peer.ID, rid cid.Cid) (thread.Record, error)

	// Updates returns a read-only channel of updates.
	Updates() <-chan *Record

	// Pull for new records from the given thread.
	// Logs owned by this host are traversed locally.
	// Remotely addressed logs are pulled from the network.
	Pull(ctx context.Context, id thread.ID) error

	// GetLogs returns info about the logs in the given thread.
	GetLogs(id thread.ID) []thread.LogInfo

	// Delete the given thread.
	Delete(ctx context.Context, id thread.ID) error
}

// Record is used to wrap a thread.Record with its thread and log context.
type Record struct {
	ThreadID thread.ID
	LogID    peer.ID
	Record   thread.Record
}
