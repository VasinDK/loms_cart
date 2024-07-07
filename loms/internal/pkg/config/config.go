package config

import "os"

var (
	Port                  = "50051"
	HttpPort              = "8085"
	Host                  = "localhost"
	DBConnect             = "postgres://admin_loms:password@localhost:5432/loms" // в docker postgres вместо localhost
	TraceEndpointURL      = "http://jaeger:4318"
	DeploymentEnvironment = "development"
)

// Config - конфигурация приложения
type Config struct {
	port                  string // grpc порт приложения
	httpPort              string // Http порт приложения
	host                  string // Хост
	dbConnect             string // Строка подключения
	TraceEndpointURL      string // Адрес куда отправляет данные трейс экспортер
	DeploymentEnvironment string // Среда развертывания
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

	if len(os.Getenv("DB_CONNECTION")) > 0 {
		DBConnect = os.Getenv("DB_CONNECTION")
	}

	if len(os.Getenv("TRACE_END_POINT_URL")) > 0 {
		TraceEndpointURL = os.Getenv("TRACE_END_POINT_URL")
	}

	if len(os.Getenv("DEPLOYMENT_ENVIRONMENT")) > 0 {
		DeploymentEnvironment = os.Getenv("DEPLOYMENT_ENVIRONMENT")
	}

	return &Config{
		port:                  Port,
		httpPort:              HttpPort,
		host:                  Host,
		dbConnect:             DBConnect,
		TraceEndpointURL:      TraceEndpointURL,
		DeploymentEnvironment: DeploymentEnvironment,
	}
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
	return c.TraceEndpointURL
}

// GetDeploymentEnvironment - среда развертывания
func (c *Config) GetDeploymentEnvironment() string {
	return c.DeploymentEnvironment
}
