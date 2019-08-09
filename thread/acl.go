package thread

import (
	"github.com/ipfs/go-ipld-format"
)

type Role int

const (
	NOACCESS Role = iota
	FOLLOW
	READ
	WRITE
	DELETE
)

type Roles interface {
	All() Role
	Actors() map[ID]Role
}

type ACL interface {
	format.Node

	Roles
	Docs() map[ID]Roles
}
