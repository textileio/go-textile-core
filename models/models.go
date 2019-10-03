// Built-in models projected from events.
package models

import (
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
)

// Metadata holds info pertaining to event retention.
type Metadata interface {
	// The max age of an event after which it can be discarded.
	MaxAge() time.Duration

	// The max count of events in a thread after which the oldest can be discarded.
	MaxCount() int
}

// Role represents a peer access role.
type Role int

const (
	// NOACCESS is granted.
	NOACCESS Role = iota
	// FOLLOW access is granted.
	FOLLOW
	// FOLLOW and READ access is granted.
	READ
	// WRITE access is granted.
	WRITE
	// DELETE access is granted.
	DELETE
)

// Roles defines the default and peer-based access roles for a thread / doc.
type Roles interface {
	// Default holds a default Role for all peers.
	Default() Role

	// Peers holds Roles for specific peers.
	Peers() map[peer.ID]Role
}

// ACL defines the access roles for a thread and its docs.
type ACL interface {
	Roles
	Docs() map[string]Roles
}
