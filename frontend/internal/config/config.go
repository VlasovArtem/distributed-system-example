package config

type Config struct {
	HTTP        HTTP
	Books       Books
	Authors     Authors
}

type HTTP struct {
	Port int
}

type Books struct {
	RPC string
}

type Authors struct {
	RPC string
}
