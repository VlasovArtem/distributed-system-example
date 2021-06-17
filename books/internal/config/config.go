package config

type Config struct {
	HTTP        HTTP
	AuthorsHTTP AuthorsHTTP
}

type HTTP struct {
	Port    int
}

type AuthorsHTTP struct {
	URL string
}
