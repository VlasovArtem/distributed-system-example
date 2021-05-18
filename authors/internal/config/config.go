package config

type Config struct {
	HTTP HTTP
	RPC  RPC
}

type HTTP struct {
	Port    int
	Enabled bool
}

type RPC struct {
	TCPPort int
	Enabled bool
}
