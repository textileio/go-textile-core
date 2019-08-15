package threadstore

import (
	"fmt"

	"github.com/ipfs/go-cid"
	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-textile-core/thread"
)

var ErrNotFound = fmt.Errorf("item not found")

type Threadstore interface {
	// Ideally this would be Close, but it overlaps with host.Host's Close
	Shutdown() error

	LogKeyBook
	LogAddrBook
	LogHeadBook

	ThreadInfo(thread.ID) thread.Info

	Threads() thread.IDSlice
}

type ThreadMetadata interface {
	Get(t thread.ID, key string) (interface{}, error)
	Put(t thread.ID, key string, val interface{}) error
}

type LogKeyBook interface {
	LogPubKey(thread.ID, peer.ID) ic.PubKey
	AddLogPubKey(thread.ID, peer.ID, ic.PubKey) error

	LogPrivKey(thread.ID, peer.ID) ic.PrivKey
	AddLogPrivKey(thread.ID, peer.ID, ic.PrivKey) error

	LogReadKey(thread.ID, peer.ID) []byte
	AddLogReadKey(thread.ID, peer.ID, []byte) error

	LogFollowKey(thread.ID, peer.ID) []byte
	AddLogFollowKey(thread.ID, peer.ID, []byte) error

	LogsWithKeys(thread.ID) peer.IDSlice

	ThreadsFromKeys() thread.IDSlice
}

type LogAddrBook interface {
	AddLogAddr(thread.ID, peer.ID, ma.Multiaddr)
	AddLogAddrs(thread.ID, peer.ID, []ma.Multiaddr)

	SetLogAddr(thread.ID, peer.ID, ma.Multiaddr)
	SetLogAddrs(thread.ID, peer.ID, []ma.Multiaddr)

	UpdateLogAddrs(t thread.ID, l peer.ID, oldAddr ma.Multiaddr, newAddr ma.Multiaddr)
	LogAddrs(thread.ID, peer.ID) []ma.Multiaddr

	ClearLogAddrs(thread.ID, peer.ID)

	LogsWithAddrs(thread.ID) peer.IDSlice

	ThreadsFromAddrs() thread.IDSlice
}

type LogHeadBook interface {
	AddLogHead(thread.ID, peer.ID, cid.Cid)
	AddLogHeads(thread.ID, peer.ID, []cid.Cid)

	SetLogHead(thread.ID, peer.ID, cid.Cid)
	SetLogHeads(thread.ID, peer.ID, []cid.Cid)

	LogHeads(thread.ID, peer.ID) []cid.Cid

	ClearLogHeads(thread.ID, peer.ID)
}

// for the wire, move to pb
type Snapshot struct {
	Logs map[peer.ID]LogSnapshot
}

type LogSnapshot struct {
	PubKey  ic.PubKey
	PrivKey ic.PrivKey
	Addrs   []ma.Multiaddr
	Heads   []cid.Cid
}
