package threads

import (
	"io"

	"github.com/ipfs/go-cid"
	"github.com/textileio/go-textile-core/threads"
)

type Threadstore interface {
	io.Closer

	Put(threads.Thread) (threads.Thread, error)
	Get(cid.Cid) threads.Thread

	Threads() []threads.Thread
	ThreadInfo(cid.Cid)
}
