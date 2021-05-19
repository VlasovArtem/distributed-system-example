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
	Enabled bool
}

type MessageQueue struct {
	URL string
}
