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
	dbConnect0            string `env:"DB_CONNECTION_0"`        // Строка подключения к postgres-shard0
	dbConnect1            string `env:"DB_CONNECTION_1"`        // Строка подключения к postgres-shard1
	traceEndpointURL      string `env:"TRACE_END_POINT_URL"`    // Адрес куда отправляет данные трейс экспортер
	deploymentEnvironment string `env:"DEPLOYMENT_ENVIRONMENT"` // Среда развертывания
	brokers               string `env:"BOOTSTRAP_SERVER"`       // брокеры
	sequenceShift         string `env:"SEQUENCE_SHIFT"`         // Сдвиг id последовательности
	mainShard             string `env:"MAIN_SHARD"`             // Главный Шард БД
}

// New - создает экземпляр конфига
func New() *Config {
	Config := &Config{
		port:                  getEnv("PORT", "50051"),
		httpPort:              getEnv("HTTP_PORT", "8085"),
		host:                  getEnv("HOST", "localhost"),
		dbConnect0:            getEnv("DB_CONNECTION_0", "postgres://admin_loms:password@localhost:5432/loms"),
		dbConnect1:            getEnv("DB_CONNECTION_1", "postgres://admin_loms:password@localhost:5433/loms"),
		traceEndpointURL:      getEnv("TRACE_END_POINT_URL", "http://localhost:4318"),
		deploymentEnvironment: getEnv("DEPLOYMENT_ENVIRONMENT", "development"),
		brokers:               getEnv("BOOTSTRAP_SERVER", "localhost:9092"),
		sequenceShift:         getEnv("SEQUENCE_SHIFT", "1000"),
		mainShard:             getEnv("MAIN_SHARD", "0"),
	}

	return Config
}

// getEnv - задаем ENV либо значение по умолчанию
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

// GetDBConnect - получает слайс строк для подключения к бд
func (c *Config) GetDBConnect() *[]string {
	return &[]string{
		c.dbConnect0,
		c.dbConnect1,
	}
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

// GetSequenceShift - Сдвиг id последовательности
func (c *Config) GetSequenceShift() string {
	return c.sequenceShift
}

// GetMainShard - Главный шард БД
func (c *Config) GetMainShard() string {
	return c.mainShard
}
