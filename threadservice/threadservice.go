package threadservice

import (
	format "github.com/ipfs/go-ipld-format"
	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/textileio/go-textile-core/thread"
	"github.com/textileio/go-textile-core/threadstore"
)

type Threadservice interface {
	// Thread actor's private key
	ic.PrivKey

	// Host provides a network identity
	host.Host

	// DAGService provides a DAG API for reading and writing thread logs
	format.DAGService

	// Threadstore persists thread log details
	threadstore.Threadstore

	// Put an event to existing threads (creates a new thread if no threads are given)
	Put(event thread.Event, t ...thread.ID)

	// Events paginates ordered thread log events
	Events(offset, limit int, t thread.ID)

	// Leave a thread
	Leave(thread.ID)

	// Delete a thread (requires ACL check)
	Delete(thread.ID)
}
