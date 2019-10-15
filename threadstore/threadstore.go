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

	Threads() (thread.IDSlice, error)
	ThreadInfo(thread.ID) (thread.Info, error)

	AddLog(thread.ID, thread.LogInfo) error
	LogInfo(thread.ID, peer.ID) (thread.LogInfo, error)
}

// ThreadMetadata stores local thread metadata like name.
type ThreadMetadata interface {
	GetInt64(t thread.ID, key string) (*int64, error)
	PutInt64(t thread.ID, key string, val int64) error
	GetString(t thread.ID, key string) (*string, error)
	PutString(t thread.ID, key string, val string) error
	GetBytes(t thread.ID, key string) (*[]byte, error)
	PutBytes(t thread.ID, key string, val []byte) error
}

// KeyBook stores log keys.
type KeyBook interface {
	PubKey(thread.ID, peer.ID) (ic.PubKey, error)
	AddPubKey(thread.ID, peer.ID, ic.PubKey) error

	PrivKey(thread.ID, peer.ID) (ic.PrivKey, error)
	AddPrivKey(thread.ID, peer.ID, ic.PrivKey) error

	ReadKey(thread.ID, peer.ID) ([]byte, error)
	AddReadKey(thread.ID, peer.ID, []byte) error

	FollowKey(thread.ID, peer.ID) ([]byte, error)
	AddFollowKey(thread.ID, peer.ID, []byte) error

	LogsWithKeys(thread.ID) (peer.IDSlice, error)

	ThreadsFromKeys() (thread.IDSlice, error)
}

// AddrBook stores log addresses.
type AddrBook interface {
	AddAddr(thread.ID, peer.ID, ma.Multiaddr, time.Duration) error
	AddAddrs(thread.ID, peer.ID, []ma.Multiaddr, time.Duration) error

	SetAddr(thread.ID, peer.ID, ma.Multiaddr, time.Duration) error
	SetAddrs(thread.ID, peer.ID, []ma.Multiaddr, time.Duration) error

	UpdateAddrs(t thread.ID, id peer.ID, oldTTL time.Duration, newTTL time.Duration) error
	Addrs(thread.ID, peer.ID) ([]ma.Multiaddr, error)

	AddrStream(context.Context, thread.ID, peer.ID) (<-chan ma.Multiaddr, error)

	ClearAddrs(thread.ID, peer.ID) error

	LogsWithAddrs(thread.ID) (peer.IDSlice, error)

	ThreadsFromAddrs() (thread.IDSlice, error)
}

// HeadBook stores log heads.
type HeadBook interface {
	AddHead(thread.ID, peer.ID, cid.Cid) error
	AddHeads(thread.ID, peer.ID, []cid.Cid) error

	SetHead(thread.ID, peer.ID, cid.Cid) error
	SetHeads(thread.ID, peer.ID, []cid.Cid) error

	Heads(thread.ID, peer.ID) ([]cid.Cid, error)

	ClearHeads(thread.ID, peer.ID) error
}
