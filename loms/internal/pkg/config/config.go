package config

import (
	"os"
	"strings"
)

// Config - конфигурация приложения
type Config struct {
	port                  string `env:"PORT"`                   // grpc порт приложения
	httpPort              string `env:"HTTP_PORT"`              // Http порт приложения
	host                  string `env:"HOST"`                   // Хост
	dbConnect             string `env:"DB_CONNECTION"`          // Строка подключения
	traceEndpointURL      string `env:"TRACE_END_POINT_URL"`    // Адрес куда отправляет данные трейс экспортер
	deploymentEnvironment string `env:"DEPLOYMENT_ENVIRONMENT"` // Среда развертывания
	brokers               string `env:"BOOTSTRAP_SERVER"`       // брокеры
}

// New - создает экземпляр конфига
func New() *Config {
	Config := &Config{
		port:                  getEnv("PORT", "50051"),
		httpPort:              getEnv("HTTP_PORT", "8085"),
		host:                  getEnv("HOST", "localhost"),
		dbConnect:             getEnv("DB_CONNECTION", "postgres://admin_loms:password@localhost:5432/loms"),
		traceEndpointURL:      getEnv("TRACE_END_POINT_URL", "http://localhost:4318"),
		deploymentEnvironment: getEnv("DEPLOYMENT_ENVIRONMENT", "development"),
		brokers:               getEnv("BOOTSTRAP_SERVER", "localhost:9092"),
	}

	return Config
}

func getEnv(key, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return defaultValue
}

// GetPort - получает grpc порт
func (c *Config) GetPort() string {
	return c.port
}

// GetHttpPort - получает http порт
func (c *Config) GetHttpPort() string {
	return c.httpPort
}

// GetHost - получает хост
func (c *Config) GetHost() string {
	return c.host
}

// GetDBConnect - получает строку подключения к бд
func (c *Config) GetDBConnect() string {
	return c.dbConnect
}

// GetTraceEndpointURL - адрес куда отправляет данные трейс экспортер
func (c *Config) GetTraceEndpointURL() string {
	return c.traceEndpointURL
}

// GetDeploymentEnvironment - среда развертывания
func (c *Config) GetDeploymentEnvironment() string {
	return c.deploymentEnvironment
}

// GetBrokers - брокеры сообщений
func (c *Config) GetBrokers() *[]string {
	brokers := strings.Split(c.brokers, ",")
	return &brokers
}
