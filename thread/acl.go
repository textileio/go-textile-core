package thread

import "github.com/libp2p/go-libp2p-core/peer"

type Role int

const (
	NOACCESS Role = iota
	FOLLOW
	READ
	WRITE
	DELETE
)

type Roles interface {
	Default() Role
	Peers() map[peer.ID]Role
}

type ACL interface {
	Roles
	Docs() map[ID]Roles
}
