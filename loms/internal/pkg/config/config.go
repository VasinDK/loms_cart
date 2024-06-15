package config

import "os"

var (
	Port = "50051"
)

// Config - конфигурация приложения
type Config struct {
	Port string // Port - порт приложения
}

// New - создает экземпляр конфига
func New() *Config {
	if len(os.Getenv("PORT")) > 0 {
		Port = os.Getenv("PORT")
	}

	return &Config{
		Port: Port,
	}
}

// GetPort - получает порт
func (c *Config) GetPort() string {
	return c.Port
}
