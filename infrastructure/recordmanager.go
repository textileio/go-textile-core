package infrastructure

import (
	datastore "github.com/ipfs/go-datastore"
	broadcast "github.com/textileio/go-textile-core/broadcast"
)

// RecordManager is used to write indexed items to database records.
type RecordManager interface {
	datastore.Datastore

	// Constructs and returns an ORM byte array from a given item object.
	ToRecord(item IndexedItem) (key datastore.Key, record []byte, err error)
	// Constructs and returns a sequenced item object, from given ORM byte array.
	FromRecord(key datastore.Key, record []byte) (item IndexedItem, err error)
	// Notifications returns a broadcast channel for relying added records.
	// This is used by downstream views to subscribe the persisted events.
	Notifications() *broadcast.Broadcaster
}

// ACIDRecordManager can write tracking records and indexed items in an atomic transaction.
type ACIDRecordManager interface {
	RecordManager

	NewTransaction(readOnly bool) (datastore.Txn, error)
}
