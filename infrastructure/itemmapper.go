package infrastructure

import (
	"github.com/textileio/go-textile-core/thread"
)

// ItemMapper maps between indexed items and (thread) events.
type ItemMapper interface {
	// Constructs and returns an indexed item for a given thread event.
	ItemFromEvent(e thread.Event) (IndexedItem, error)
	// Constructs and returns a thread event for given indexed item.
	// @note: Implementations will likely have to fetch the event using its Cid?
	EventFromItem(i IndexedItem) (thread.Event, error)
}
