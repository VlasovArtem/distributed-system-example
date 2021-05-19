package config

type Config struct {
	HTTP        HTTP
	Books       Books
	Authors     Authors
	RPCEnabled  bool
}

type HTTP struct {
	Port int
}

type Books struct {
	URL string
	RPC string
}

type Authors struct {
	URL string
	RPC string
}
