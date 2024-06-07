package config

import "os"

var (
	Port         = "8082"
	TokenStore   = "testtoken"
	AddressStore = "http://route256.pavl.uk:8080/get_product"
)

type Config struct {
	Port         string
	TokenStore   string
	AddressStore string
}

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

	return &Config{
		Port:         ":" + Port,
		TokenStore:   TokenStore,
		AddressStore: AddressStore,
	}
}

func (c *Config) GetPort() string {
	return c.Port
}

func (c *Config) GetTokenStore() string {
	return c.TokenStore
}

func (c *Config) GetAddressStore() string {
	return c.AddressStore
}
