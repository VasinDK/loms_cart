package config

import (
	"os"
	"strings"
)

// Config - конфигурация приложения
type Config struct {
	ConsumerGroupName string    `env:"CONSUMER_GROUP"`   // Консюмер группа
	Topic             *[]string `env:"TOPIC"`            // Топик продюсера
	Brokers           *[]string `env:"BOOTSTRAP_SERVER"` // хост:порт кафка брокера
}

// New - создает экземпляр конфига
func New() *Config {
	Config := &Config{
		ConsumerGroupName: getEnv("CONSUMER_GROUP", "CG_BASE"),
		Topic:             getEnvByStr("TOPIC", "loms.order-events"),
		Brokers:           getEnvByStr("BOOTSTRAP_SERVER", "localhost:9092"),
	}

	return Config
}

func getEnv(key, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaultValue
}

func getEnvByStr(key, defaultValue string) *[]string {
	if v, ok := os.LookupEnv(key); ok {
		res := strings.Split(v, ",")

		return &res
	}
	return &[]string{defaultValue}
}

// GetBrokers -
func (c *Config) GetBrokers() *[]string {
	return c.Brokers
}

// GetGroupID -
func (c *Config) GetGroupID() string {
	return c.ConsumerGroupName
}

// GetTopics -
func (c *Config) GetTopics() *[]string {
	return c.Topic
}
