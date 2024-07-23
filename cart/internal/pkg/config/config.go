package config

import (
	"context"
	"os"
	"route256/cart/internal/pkg/logger"
	"strconv"
)

// Config - конфигурация приложения
type Config struct {
	Port                  string `env:"PORT"`                   // Port - порт приложения
	TokenStore            string `env:"TOKEN_STORE"`            // TokenStore - токен для стороннего хранилища
	AddressStore          string `env:"ADDRESS_STORE"`          // AddressStore - адрес стороннего хранилища
	AddressStoreLoms      string `env:"ADDRESS_STORE_LOMS"`     // AddressStoreLoms - адрес Loms хранилища
	PostLoms              string `env:"PORT_LOMS"`              // Port grpc loms
	TimeGraceShutdown     int64  `env:"TIME_GRACE_SHUTDOWN"`    // Время для плавного завершения работы сервера
	TraceEndpointURL      string `env:"TRACE_END_POINT_URL"`    // Адрес куда отправляет данные трейс экспортер
	DeploymentEnvironment string `env:"DEPLOYMENT_ENVIRONMENT"` // Среда развертывания
	SizeBufferCache       int64  `env:"SIZE_BUFFER_CACHE"`      // Размер буфера кеша
}

// New - создает экземпляр конфига
func New() *Config {
	Config := &Config{
		Port:                  getEnvStr("PORT", "8082"),
		TokenStore:            getEnvStr("TOKEN_STORE", "testtoken"),
		AddressStore:          getEnvStr("ADDRESS_STORE", "http://route256.pavl.uk:8080/get_product"),
		AddressStoreLoms:      getEnvStr("ADDRESS_STORE_LOMS", "localhost"),
		PostLoms:              getEnvStr("PORT_LOMS", "50051"),
		TimeGraceShutdown:     getEnvInt64("TIME_GRACE_SHUTDOWN", 5),
		TraceEndpointURL:      getEnvStr("TRACE_END_POINT_URL", "http://localhost:4318"),
		DeploymentEnvironment: getEnvStr("DEPLOYMENT_ENVIRONMENT", "development"),
		SizeBufferCache:       getEnvInt64("SIZE_BUFFER_CACHE", 5),
	}

	return Config
}

func getEnvStr(key, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaultValue
}

func getEnvInt64(key string, defaultValue int64) int64 {
	if v, ok := os.LookupEnv(key); ok {
		NewValue, err := strconv.Atoi(v)
		if err != nil {
			logger.Errorw(context.Background(), "getEnvInt env", err.Error())
			return defaultValue
		}
		return int64(NewValue)
	}
	return defaultValue
}

// GetPort - получает порт
func (c *Config) GetPort() string {
	return c.Port
}

// GetPortLoms - получает порт
func (c *Config) GetPortLoms() string {
	return c.PostLoms
}

// GetTokenStore - получает токен
func (c *Config) GetTokenStore() string {
	return c.TokenStore
}

// GetAddressStore - получает адрес хранилища
func (c *Config) GetAddressStore() string {
	return c.AddressStore
}

// GetAddressStoreLoms - получает адрес хранилища Loms
func (c *Config) GetAddressStoreLoms() string {
	return c.AddressStoreLoms
}

// GetTimeGraceShutdown - получает время необходимое для GraceShutdown
func (c *Config) GetTimeGraceShutdown() int64 {
	return c.TimeGraceShutdown
}

// GetTraceEndpointURL - адрес куда отправляет данные трейс экспортер
func (c *Config) GetTraceEndpointURL() string {
	return c.TraceEndpointURL
}

// GetDeploymentEnvironment - среда развертывания
func (c *Config) GetDeploymentEnvironment() string {
	return c.DeploymentEnvironment
}

// GetSizeBufferCache - возвращает размер буфера
func (c *Config) GetSizeBufferCache() int64 {
	return c.SizeBufferCache
}
