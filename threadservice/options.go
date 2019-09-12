package threadservice

import (
	"github.com/textileio/go-textile-core/thread"
)

// PutOpt is an instance helper for creating options
var PutOpt PutOption

// PutOption is used to create putSettings
type PutOption func(*putSettings)

// Thread sets the target thread for a put operation
func (PutOption) Thread(val thread.ID) PutOption {
	return func(settings *putSettings) {
		settings.Thread = val
	}
}

// putSettings holds values used for a a put operation
type putSettings struct {
	Thread thread.ID
}

// PutOptions returns put settings from options
func PutOptions(opts ...PutOption) *putSettings {
	options := &putSettings{
		Thread: thread.NewIDV1(thread.AccessControlled, 16),
	}

	for _, opt := range opts {
		opt(options)
	}
	return options
}
