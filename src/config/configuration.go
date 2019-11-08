package config

type DatabaseConfiguration struct {
	ConnectionUri string
}

type ServerConfiguration struct {
	Port int
}

type Configuration struct {
	Server ServerConfiguration
	Database DatabaseConfiguration
}