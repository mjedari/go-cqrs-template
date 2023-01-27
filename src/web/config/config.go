package config

var Config Configuration

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
}

type HTTPClient struct {
	Timeout int64
}

type Redis struct {
	Host string
	Port string
	User string
	Pass string
}

type Configuration struct {
	Server     Server
	Database   Database
	HTTPClient HTTPClient
	Redis      Redis
}
