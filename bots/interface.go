// Package shared contains shared data between the host and plugins.
package bots

// BotConfig is pulled from the bot config file
type BotConfig struct {
	BotName        string
	BotID          string
	ReleaseVersion int
	ReleaseHash    string
	Params         map[string]string
}

// SharedConfig contain all the services and config passed by the host node
type SharedConfig struct {
	Store  BotStore
	Ipfs   IpfsHandler
	Params map[string]string
}

// Response is the response for each request to a Bot
type Response struct {
	Status      int32
	Body        []byte
	ContentType string
}

// Botstore is an interface that should be provided by the Cafe to get/set to a storage backend
type BotStore interface {
	// TODO: stored data []byte might be better as json or json string objects?
	Set(key string, data []byte) (ok bool, err error)
	Get(key string) (data []byte, version int32, err error)
	Delete(key string) (ok bool, err error)
	// TODO: how to manage cleanup? e.g. expired links should be removed occasionally
}

// IpfsHandler is an interface to the gateway method to fetch + decrypt content
type IpfsHandler interface {
	Get(path string, key string) (data []byte, err error)
	Add(data []byte, encrypt bool) (hash string, key string, err error)
}

type Botservice interface {
	Post(data []byte, body []byte, shared SharedConfig) (Response, error)
	Get(data []byte, shared SharedConfig) (Response, error)
	Put(data []byte, body []byte, shared SharedConfig) (Response, error)
	Delete(data []byte, shared SharedConfig) (Response, error)
}
