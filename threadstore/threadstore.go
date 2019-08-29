package threadstore

import (
	"context"
	"fmt"
	"time"

	"github.com/ipfs/go-cid"
	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-textile-core/thread"
)

// ErrNotFound is and error used to indicate an item is not found
var ErrNotFound = fmt.Errorf("item not found")

// Threadstore stores log keys, addresses, heads and thread meta data
type Threadstore interface {
	Shutdown() error

	ThreadMetadata
	LogKeyBook
	LogAddrBook
	LogHeadBook

	ThreadInfo(thread.ID) thread.Info

	Threads() thread.IDSlice
}

// ThreadMetadata
type ThreadMetadata interface {
	Get(t thread.ID, key string) (interface{}, error)
	Put(t thread.ID, key string, val interface{}) error
}

// LogKeyBook stores log keys
type LogKeyBook interface {
	PubKey(thread.ID, peer.ID) ic.PubKey
	AddPubKey(thread.ID, peer.ID, ic.PubKey) error

	PrivKey(thread.ID, peer.ID) ic.PrivKey
	AddPrivKey(thread.ID, peer.ID, ic.PrivKey) error

	ReadKey(thread.ID, peer.ID) []byte
	AddReadKey(thread.ID, peer.ID, []byte) error

	FollowKey(thread.ID, peer.ID) []byte
	AddFollowKey(thread.ID, peer.ID, []byte) error

	LogsWithKeys(thread.ID) peer.IDSlice

	ThreadsFromKeys() thread.IDSlice
}

// LogAddrBook stores log addresses
type LogAddrBook interface {
	AddAddr(thread.ID, peer.ID, ma.Multiaddr, time.Duration)
	AddAddrs(thread.ID, peer.ID, []ma.Multiaddr, time.Duration)

	SetAddr(thread.ID, peer.ID, ma.Multiaddr, time.Duration)
	SetAddrs(thread.ID, peer.ID, []ma.Multiaddr, time.Duration)

	UpdateAddrs(t thread.ID, p peer.ID, oldTTL time.Duration, newTTL time.Duration)
	Addrs(thread.ID, peer.ID) []ma.Multiaddr

	AddrStream(context.Context, thread.ID, peer.ID) <-chan ma.Multiaddr

	ClearAddrs(thread.ID, peer.ID)

	LogsWithAddrs(thread.ID) peer.IDSlice

	ThreadsFromAddrs() thread.IDSlice
}

// LogHeadBook stores log heads
type LogHeadBook interface {
	AddHead(thread.ID, peer.ID, cid.Cid)
	AddHeads(thread.ID, peer.ID, []cid.Cid)

	SetHead(thread.ID, peer.ID, cid.Cid)
	SetHeads(thread.ID, peer.ID, []cid.Cid)

	Heads(thread.ID, peer.ID) []cid.Cid

	ClearHeads(thread.ID, peer.ID)
}

// for the wire, move to pb
type Snapshot struct {
	Logs map[peer.ID]LogSnapshot
}

// for the wire, move to pb
type LogSnapshot struct {
	PubKey  ic.PubKey
	PrivKey ic.PrivKey
	Addrs   []ma.Multiaddr
	Heads   []cid.Cid
}
