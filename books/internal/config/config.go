package config

type Config struct {
	RPC RPC
	AuthorsRPC AuthorsRPC
}

type AuthorsRPC struct {
	URL string
}

type RPC struct {
	TCPPort int
}
