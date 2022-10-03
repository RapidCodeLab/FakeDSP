package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type ExtendedServerConfig struct {
	HttpServerConfig
	MetricsListenAddr string `env:"METRICS_LISTEN_ADDR" env-default:":9090"`
}

type HttpServerConfig struct {
	ListenNetwork string `env:"LISTEN_NETWORK" env-default:"tcp4"`
	ListenAddr    string `env:"LISTEN_ADDR" env-default:":8080"`
}

func (c *HttpServerConfig) GetListenAddr() string {
	return c.ListenAddr
}

func (c *HttpServerConfig) GetListenNetwork() string {
	return c.ListenNetwork
}

func GetHTTPServerConfig() (*HttpServerConfig, error) {
	cfg := &HttpServerConfig{}
	return cfg, ParseENV(cfg)
}

func GetExtendedConfig() (*ExtendedServerConfig, error) {
	cfg := &ExtendedServerConfig{}
	return cfg, ParseENV(cfg)
}

func ParseENV(cfg interface{}) error {
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		help, err := cleanenv.GetDescription(cfg, nil)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s | %s", help, err.Error())
	}

	return nil

}
