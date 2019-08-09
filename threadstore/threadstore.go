package threadstore

import (
	"io"

	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-textile-core/thread"
)

type Threadstore interface {
	io.Closer

	LogKeyBook
	LogAddrBook

	ThreadInfo(thread.ID) thread.Info

	Threads() thread.IDSlice
}

type ThreadMetadata interface {
	Get(t thread.ID, key string) (interface{}, error)
	Put(t thread.ID, key string, val interface{}) error
}

type LogKeyBook interface {
	PubKey(thread.ID, peer.ID) ic.PubKey
	AddPubKey(thread.ID, ic.PubKey) error

	PrivKey(thread.ID, peer.ID) ic.PrivKey
	AddPrivKey(thread.ID, ic.PrivKey) error

	ReadKey(thread.ID, peer.ID) []byte
	AddReadKey(thread.ID, peer.ID, []byte) error

	FollowKey(thread.ID, peer.ID) []byte
	AddFollowKey(thread.ID, peer.ID, []byte) error

	LogsWithKeys(thread.ID) peer.IDSlice

	ThreadsFromKeys() thread.IDSlice
}

type LogAddrBook interface {
	AddAddr(t thread.ID, l peer.ID, addr ma.Multiaddr)
	AddAddrs(t thread.ID, l peer.ID, addrs []ma.Multiaddr)

	SetAddr(t thread.ID, l peer.ID, addr ma.Multiaddr)
	SetAddrs(t thread.ID, l peer.ID, addrs []ma.Multiaddr)

	UpdateAddrs(t thread.ID, l peer.ID, oldAddr ma.Multiaddr, newAddr ma.Multiaddr)
	Addrs(t thread.ID, l peer.ID) []ma.Multiaddr

	ClearAddrs(t thread.ID, l peer.ID)

	LogsWithAddrs(t thread.ID) peer.IDSlice

	ThreadsFromAddrs() thread.IDSlice
}
