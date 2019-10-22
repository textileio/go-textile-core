package threadservice

import (
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-textile-core/crypto"
	"github.com/textileio/go-textile-core/thread"
)

// AddOpt is an instance helper for creating add options.
var AddOpt AddOption

// AddOption is used to create AddSettings.
type AddOption func(*AddSettings)

// ThreadID sets the target thread for an add operation.
func (AddOption) ThreadID(val thread.ID) AddOption {
	return func(settings *AddSettings) {
		settings.ThreadID = val
	}
}

// Time sets wall-clock time associated for an add operation.
func (AddOption) Time(val time.Time) AddOption {
	return func(settings *AddSettings) {
		settings.Time = val
	}
}

// Key sets the read encryption key used for an add operation.
// If both KeyLog and this option are absent, the target log's read key is used.
// This option takes precedence over KeyLogID.
func (AddOption) Key(val crypto.EncryptionKey) AddOption {
	return func(settings *AddSettings) {
		settings.Key = val
	}
}

// KeyLog sets the read encryption key used for an add operation
// to the read key of the given log ID.
// If both Key and this option are absent, the target log's read key is used.
func (AddOption) KeyLog(val peer.ID) AddOption {
	return func(settings *AddSettings) {
		settings.KeyLog = val
	}
}

// Addrs holds additional addresses to notify with an add operation.
func (AddOption) Addrs(val []ma.Multiaddr) AddOption {
	return func(settings *AddSettings) {
		settings.Addrs = val
	}
}

// AddSettings holds values used for an add operation.
type AddSettings struct {
	ThreadID thread.ID
	Time     time.Time
	Key      crypto.EncryptionKey
	KeyLog   peer.ID
	Addrs    []ma.Multiaddr
}

// AddOptions returns add settings from options.
func AddOptions(opts ...AddOption) *AddSettings {
	options := &AddSettings{
		ThreadID: thread.NewIDV1(thread.AccessControlled, 16),
		Time:     time.Now(),
	}

	for _, opt := range opts {
		opt(options)
	}
	return options
}

// PutOpt is an instance helper for creating put options.
var PutOpt PutOption

// PutOption is used to create PutSettings.
type PutOption func(*PutSettings)

// ThreadID sets the target thread for a put operation.
func (PutOption) ThreadID(val thread.ID) PutOption {
	return func(settings *PutSettings) {
		settings.ThreadID = val
	}
}

// LogID sets the target log for a put operation.
func (PutOption) LogID(val peer.ID) PutOption {
	return func(settings *PutSettings) {
		settings.LogID = val
	}
}

// PutSettings holds values used for a put operation.
type PutSettings struct {
	ThreadID thread.ID
	LogID    peer.ID
}

// PutOptions returns put settings from options.
func PutOptions(opts ...PutOption) *PutSettings {
	options := &PutSettings{
		ThreadID: thread.NewIDV1(thread.AccessControlled, 16),
	}

	for _, opt := range opts {
		opt(options)
	}
	return options
}

// SubOpt is an instance helper for creating subscription options.
var SubOpt SubOption

// SubOption is used to create SubSettings.
type SubOption func(*SubSettings)

// ThreadID restricts the subscription to the given thread.
// Use this option multiple times to build up a list of threads
// to subscribe to.
func (SubOption) ThreadID(val thread.ID) SubOption {
	return func(settings *SubSettings) {
		settings.ThreadIDs = append(settings.ThreadIDs, val)
	}
}

// SubSettings holds values used for a subscribe operation.
type SubSettings struct {
	ThreadIDs []thread.ID
}

// SubOptions returns subscription settings from options.
func SubOptions(opts ...SubOption) *SubSettings {
	options := &SubSettings{}

	for _, opt := range opts {
		opt(options)
	}
	return options
}
