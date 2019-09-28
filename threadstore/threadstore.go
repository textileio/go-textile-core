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

// ErrNotFound is and error used to indicate an item is not found.
var ErrNotFound = fmt.Errorf("item not found")

// Threadstore stores log keys, addresses, heads and thread meta data.
type Threadstore interface {
	Close() error

	ThreadMetadata
	KeyBook
	AddrBook
	HeadBook

	Threads() thread.IDSlice
	ThreadInfo(thread.ID) thread.Info

	AddLog(thread.ID, thread.LogInfo) error
	LogInfo(thread.ID, peer.ID) thread.LogInfo
}

// ThreadMetadata stores local thread metadata like name.
type ThreadMetadata interface {
	GetMeta(t thread.ID, key string) (interface{}, error)
	PutMeta(t thread.ID, key string, val interface{}) error
}

// KeyBook stores log keys.
type KeyBook interface {
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

// AddrBook stores log addresses.
type AddrBook interface {
	AddAddr(thread.ID, peer.ID, ma.Multiaddr, time.Duration)
	AddAddrs(thread.ID, peer.ID, []ma.Multiaddr, time.Duration)

	SetAddr(thread.ID, peer.ID, ma.Multiaddr, time.Duration)
	SetAddrs(thread.ID, peer.ID, []ma.Multiaddr, time.Duration)

	UpdateAddrs(t thread.ID, id peer.ID, oldTTL time.Duration, newTTL time.Duration)
	Addrs(thread.ID, peer.ID) []ma.Multiaddr

	AddrStream(context.Context, thread.ID, peer.ID) <-chan ma.Multiaddr

	ClearAddrs(thread.ID, peer.ID)

	LogsWithAddrs(thread.ID) peer.IDSlice

	ThreadsFromAddrs() thread.IDSlice
}

// HeadBook stores log heads.
type HeadBook interface {
	AddHead(thread.ID, peer.ID, cid.Cid)
	AddHeads(thread.ID, peer.ID, []cid.Cid)

	SetHead(thread.ID, peer.ID, cid.Cid)
	SetHeads(thread.ID, peer.ID, []cid.Cid)

	Heads(thread.ID, peer.ID) []cid.Cid

	ClearHeads(thread.ID, peer.ID)
}
