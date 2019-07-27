package threads

import (
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/go-textile-wallet/account"
)

type ThreadIntent string

type Thread interface {
	format.Node

	ReaderKey() []byte
	ReplicatorKey() []byte

	Schema() Schema
	Intent() ThreadIntent
	Roles() Roles

	Heads() []cid.Cid

	GetName() string
	SetName() error

	CreateInvite() (cid.Cid, error)
	Invite(account.Account) error

	Join() error
	Leave() error

	Logs() []Log

	Write(Node) (cid.Cid, error)
	Listen() <-chan Node

	Fork() (Thread, error)
}

type Node interface {
	format.Node

	Thread() Thread
	Target() Node
	Parents() []Node
	Data() format.Node

	Signature() []byte
	Payload() []byte
}

type Log interface {
	ID() string // threadID + PeerID ?

	Author() peer.ID
	Account() account.Account

	Head() cid.Cid

	Put(Node) error
	List(/* ListOpts */)
	Remove() error
}

type Member interface {
	Account() account.Account
	Thread() cid.Cid
	Welcomed() bool
}

type Role int

const (
	NO_ACCESS Role = iota
	READ
	ANNOTATE
	WRITE
)

type Roles interface {
	Default()
	Members() map[account.Account]Role
}

type Schema interface {
	format.Node

	Name() string
	Pin() bool
	Plaintext() bool
	Mill() string
	Options() map[string]string

	JSONSchema() JSONSchema
}

type JSONSchema interface {
	format.Node
}
