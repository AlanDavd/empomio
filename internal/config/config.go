package config

import "flag"

type Config interface {
	// Server
	ServerAddress() string
	// Database
	PostgreSQLUser() string
	PostgreSQLPassword() string
	PostgreSQLName() string
	PostgreSQLHost() string
	PostgreSQLPort() int
	PostgreSQLSSLMode() string
}

type config struct {
	serverAddrs string
	// DBs
	postgreSQLUser string
	postgreSQLPassword string
	postgreSQLName string
	postgreSQLHost string
	postgreSQLPort int
	postgreSQLSSLMode string
}

func LoadConfigFlags() Config {
	serverAddrs := flag.String("address", ":8080", "Server address port")

	// Databases
	postgreSQLUser := flag.String("db_user", "empomio_user", "PostgreSQL database user")
	postgreSQLPassword := flag.String("db_pass", "empomio_password", "PostgreSQL database password")
	postgreSQLName := flag.String("db_name", "empomio", "PostgreSQL database name")
	postgreSQLHost := flag.String("db_host", "localhost", "PostgreSQL database host")
	postgreSQLPort := flag.Int("db_port", 5432, "PostgreSQL database port")
	postgreSQLSSLMode := flag.String("db_ssl_mode", "disable", "PostgreSQL database SSL mode")

	flag.Parse()
	return &config{
		serverAddrs: *serverAddrs,
		postgreSQLUser: *postgreSQLUser,
		postgreSQLPassword: *postgreSQLPassword,
		postgreSQLName: *postgreSQLName,
		postgreSQLHost: *postgreSQLHost,
		postgreSQLPort: *postgreSQLPort,
		postgreSQLSSLMode: *postgreSQLSSLMode,
	}
}

func (c *config) ServerAddress() string {
	return c.serverAddrs
}

func (c *config) PostgreSQLUser() string {
	return c.postgreSQLUser
}

func (c *config) PostgreSQLPassword() string {
	return c.postgreSQLPassword
}

func (c *config) PostgreSQLName() string {
	return c.postgreSQLName
}

func (c *config) PostgreSQLHost() string {
	return c.postgreSQLHost
}

func (c *config) PostgreSQLPort() int {
	return c.postgreSQLPort
}

func (c *config) PostgreSQLSSLMode() string {
	return c.postgreSQLSSLMode
}
