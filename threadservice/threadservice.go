package threadservice

import (
	"context"
	"net"
	"net/http"

	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/go-textile-core/thread"
	tstore "github.com/textileio/go-textile-core/threadstore"
)

// Threadservice is an API for working with threads
type Threadservice interface {
	// Threadstore persists thread log details
	tstore.Threadstore

	// Host provides a network identity
	Host() host.Host

	// DAGService provides a DAG API for reading and writing thread logs
	DAGService() format.DAGService

	// Listener is a net listener serving the threads api
	Listener() net.Listener

	// Client exposes a libp2p-router http client
	Client() *http.Client

	// Put data in existing threads (creates a new thread if no threads are given)
	Put(ctx context.Context, body format.Node, opts ...PutOption) (peer.ID, cid.Cid, error)

	// Pull paginates thread log events
	Pull(ctx context.Context, offset cid.Cid, limit int, t thread.ID, id peer.ID) ([]thread.Event, error)

	// Invite an actor to a thread
	Invite(ctx context.Context, actor peer.ID, t thread.ID) error

	// Leave a thread
	Leave(context.Context, thread.ID) error

	// Delete a thread (requires ACL check)
	Delete(context.Context, thread.ID) error
}
