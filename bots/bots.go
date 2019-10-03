// Package shared contains shared data between the host and plugins.
package bots

// BotConfig is pulled from the bot config file
type HostConfig struct {
	Name           string
	ID             string
	ReleaseVersion int
	ReleaseHash    string
	Params         map[string]string
}

// ClientConfig contain all the services and config passed by the host node
type ClientConfig struct {
	Store  Store
	Ipfs   Ipfs
	Params map[string]string
}

// Response is the response for each request to a Bot
type Response struct {
	Status      int32
	Body        []byte
	ContentType string
}

// Store is an interface that should be provided by the Cafe to get/set to a storage backend
type Store interface {
	// TODO: stored data []byte might be better as json or json string objects?
	Set(key string, data []byte) (ok bool, err error)
	Get(key string) (data []byte, version int32, err error)
	Delete(key string) (ok bool, err error)
	// TODO: how to manage cleanup? e.g. expired links should be removed occasionally
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
