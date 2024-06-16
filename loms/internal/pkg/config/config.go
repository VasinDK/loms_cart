package config

import "os"

var (
	Port     = "50051"
	HttpPort = "8085"
	Host     = "localhost"
)

// Config - конфигурация приложения
type Config struct {
	Port     string // Port - порт приложения
	HttpPort string // Http port - http порт приложения
	Host     string
}

// New - создает экземпляр конфига
func New() *Config {
	if len(os.Getenv("PORT")) > 0 {
		Port = os.Getenv("PORT")
	}

	if len(os.Getenv("HOST")) > 0 {
		Host = os.Getenv("HOST")
	}

	if len(os.Getenv("HTTP_PORT")) > 0 {
		HttpPort = os.Getenv("HTTP_PORT")
	}

	return &Config{
		Port:     Port,
		HttpPort: HttpPort,
		Host:     Host,
	}
}

// GetPort - получает порт
func (c *Config) GetPort() string {
	return c.Port
}

// GetHttpPort - получает порт
func (c *Config) GetHttpPort() string {
	return c.HttpPort
}

// GetHost - получает порт
func (c *Config) GetHost() string {
	return c.Host
}