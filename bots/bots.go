// Package bots contains shared data between the host and plugins.
package bots

import (
	"github.com/ipfs/go-datastore"
)

// HostConfig is pulled from the bot config file
type HostConfig struct {
	Name           string
	ID             string
	ReleaseVersion int
	ReleaseHash    string
	Params         map[string]string
}

// ClientConfig contain all the services and config passed by the host node
type ClientConfig struct {
	Store  DatastoreWithoutQuery
	Ipfs   Ipfs
	Params map[string]string
}

// Response is the response for each request to a Bot
type Response struct {
	Status      int32
	Body        []byte
	ContentType string
}

// Read is the read-side of the Datastore interface minus Query.
type DatastoreWithoutQuery interface {
	// See datastore Get
	Get(key datastore.Key) (value []byte, err error)
	// See datastore Has
	Has(key datastore.Key) (exists bool, err error)
	// See datastore GetSize
	GetSize(key datastore.Key) (size int, err error)
	// See datastore Put
	Put(key datastore.Key, value []byte) error
	// See datastore Delete
	Delete(key datastore.Key) error
	// See datastore Close
	Close() error

	// TODO: Implement query/result in Protobuf/rpc and go to complete Datastore
}

// Ipfs is an interface to the gateway method to fetch + decrypt content
type Ipfs interface {
	Get(path string, key string) (data []byte, err error)
	Add(data []byte, encrypt bool) (hash string, key string, err error)
}

// Service defines the methods served by any bot
type Service interface {
	Post(data []byte, body []byte, shared ClientConfig) (Response, error)
	Get(data []byte, shared ClientConfig) (Response, error)
	Put(data []byte, body []byte, shared ClientConfig) (Response, error)
	Delete(data []byte, shared ClientConfig) (Response, error)
}
