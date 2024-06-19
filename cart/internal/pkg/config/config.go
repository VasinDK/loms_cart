package config

import "os"

var (
	Port             = "8082"
	TokenStore       = "testtoken"
	AddressStore     = "http://route256.pavl.uk:8080/get_product"
	AddressStoreLoms = "localhost"
	PostLoms         = "50051"
)

// Config - конфигурация приложения
type Config struct {
	Port             string // Port - порт приложения
	TokenStore       string // TokenStore - токен для стороннего хранилища
	AddressStore     string // AddressStore - адрес стороннего хранилища
	AddressStoreLoms string // AddressStoreLoms - адрес Loms хранилища
	PostLoms         string // Port grpc loms
}

// New - создает экземпляр конфига
func New() *Config {
	if len(os.Getenv("PORT")) > 0 {
		Port = os.Getenv("PORT")
	}

	if len(os.Getenv("TOKEN_STORE")) > 0 {
		TokenStore = os.Getenv("TOKEN_STORE")
	}

	if len(os.Getenv("ADDRESS_STORE")) > 0 {
		AddressStore = os.Getenv("ADDRESS_STORE")
	}

	if len(os.Getenv("ADDRESS_STORE_LOMS")) > 0 {
		AddressStoreLoms = os.Getenv("ADDRESS_STORE_LOMS")
	}

	if len(os.Getenv("PORT_LOMS")) > 0 {
		PostLoms = os.Getenv("PORT_LOMS")
	}

	return &Config{
		Port:             Port,
		TokenStore:       TokenStore,
		AddressStore:     AddressStore,
		AddressStoreLoms: AddressStoreLoms,
		PostLoms:         PostLoms,
	}
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

// GetAddressStoreLoms - получает адрес хранилища
func (c *Config) GetAddressStoreLoms() string {
	return c.AddressStoreLoms
}
