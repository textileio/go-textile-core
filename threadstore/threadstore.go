package threadstore

import (
	"io"

	ic "github.com/libp2p/go-libp2p-core/crypto"
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
	PrivKey(thread.ID, ic.PubKey) ic.PrivKey
	AddPrivKey(thread.ID, ic.PrivKey) error

	ReadKey(thread.ID, ic.PubKey) []byte
	AddReadKey(thread.ID, ic.PubKey, []byte) error

	FollowKey(thread.ID, ic.PubKey) []byte
	AddFollowKey(thread.ID, ic.PubKey, []byte) error

	LogsWithKeys(thread.ID) []ic.PubKey
}

type LogAddrBook interface {
	AddAddr(t thread.ID, l ic.PubKey, addr ma.Multiaddr)
	AddAddrs(t thread.ID, l ic.PubKey, addrs []ma.Multiaddr)

	SetAddr(t thread.ID, l ic.PubKey, addr ma.Multiaddr)
	SetAddrs(t thread.ID, l ic.PubKey, addrs []ma.Multiaddr)

	UpdateAddrs(t thread.ID, l ic.PubKey, oldAddr ma.Multiaddr, newAddr ma.Multiaddr)
	Addrs(t thread.ID, l ic.PubKey) []ma.Multiaddr

	ClearAddrs(t thread.ID, l ic.PubKey)

	LogsWithAddrs(t thread.ID) []ic.PubKey
}
