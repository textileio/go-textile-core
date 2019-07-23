package threads

import (
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/textileio/go-textile-wallet/account"
)

type ThreadIntent string

type Thread interface {
	format.Node

	Key() []byte
	Schema() Schema
	Intent() ThreadIntent
	Public() bool
	Roles() Roles

	Heads() []cid.Cid

	GetName() string
	SetName() error

	CreateInvite() (cid.Cid, error)
	Invite(account.Account) error

	Join() error
	Leave() error
	Members()

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
