package thread

import (
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-ipld-format"
	"github.com/libp2p/go-libp2p-core/peer"
)

type Node interface {
	format.Node

	Block() format.Node
	Parents() []cid.Cid
	Signature() []byte
}

type Metadata interface {
	MaxAge() time.Duration
	MaxCount() int
	ACL() ACL
}

type ID string

type Info struct {
	ID   ID
	Logs peer.IDSlice
}

type IDSlice []ID

func (es IDSlice) Len() int           { return len(es) }
func (es IDSlice) Swap(i, j int)      { es[i], es[j] = es[j], es[i] }
func (es IDSlice) Less(i, j int) bool { return string(es[i]) < string(es[j]) }
