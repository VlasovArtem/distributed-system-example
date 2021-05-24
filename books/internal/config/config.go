package config

type Config struct {
	HTTP HTTP
	RPC  RPC
	MQ   MessageQueue
}

type HTTP struct {
	Port    int
}

type RPC struct {
	TCPPort int
}

type MessageQueue struct {
	URL string
}
