package config

import "flag"

type Config interface {
	// Server
	ServerAddress() string
}

type config struct {
	serverAddrs string
}

func LoadConfigFlags() Config {
	serverAddrs := flag.String("address", ":8080", "Server address port")
	flag.Parse()

	return &config{
		serverAddrs: *serverAddrs,
	}
}

func (c *config) ServerAddress() string {
	return c.serverAddrs
}
