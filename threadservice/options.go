package threadservice

import (
	"time"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-textile-core/crypto"
	"github.com/textileio/go-textile-core/thread"
)

// PutOpt is an instance helper for creating put options.
var PutOpt PutOption

// PutOption is used to create PutSettings.
type PutOption func(*PutSettings)

// Thread sets the target thread for a put operation.
func (PutOption) Thread(val thread.ID) PutOption {
	return func(settings *PutSettings) {
		settings.Thread = val
	}
}

// Time sets wall-clock time associated for a put operation.
func (PutOption) Time(val time.Time) PutOption {
	return func(settings *PutSettings) {
		settings.Time = val
	}
}

// Key sets the read encryption key used for a put operation.
// If no key is given, the target log's read key is used.
func (PutOption) Key(val crypto.EncryptionKey) PutOption {
	return func(settings *PutSettings) {
		settings.Key = val
	}
}

// Addrs holds additional addresses to notify with a put operation.
func (PutOption) Addrs(val []ma.Multiaddr) PutOption {
	return func(settings *PutSettings) {
		settings.Addrs = val
	}
}

// PutSettings holds values used for a a put operation.
type PutSettings struct {
	Thread thread.ID
	Time   time.Time
	Key    crypto.EncryptionKey
	Addrs  []ma.Multiaddr
}

// PutOptions returns put settings from options.
func PutOptions(opts ...PutOption) *PutSettings {
	options := &PutSettings{
		Thread: thread.NewIDV1(thread.AccessControlled, 16),
		Time:   time.Now(),
	}

	for _, opt := range opts {
		opt(options)
	}
	return options
}
