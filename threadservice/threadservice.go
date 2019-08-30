package threadservice

import (
	format "github.com/ipfs/go-ipld-format"
	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/go-textile-core/thread"
	tstore "github.com/textileio/go-textile-core/threadstore"
)

// Threadservice is an API for working with threads
type Threadservice interface {
	// Thread actor's private key
	ic.PrivKey

	// Host provides a network identity
	host.Host

	// DAGService provides a DAG API for reading and writing thread logs
	format.DAGService

	// Threadstore persists thread log details
	tstore.Threadstore

	// Put an event to existing threads (creates a new thread if no threads are given)
	Put(event thread.Event, t ...thread.ID) peer.IDSlice

	// Pull paginates ordered thread log events
	Pull(offset string, size int, t thread.ID) <-chan []thread.Event

	// Invite an actor to a thread
	Invite(actor peer.ID, t thread.ID) error

	// Leave a thread
	Leave(thread.ID) error

	// Delete a thread (requires ACL check)
	Delete(thread.ID) error
}
