package config

import (
	"github.com/caarlos0/env/v6"
)

type Conf struct {
	ServicePort           int      `env:"PORT" envDefault:"9090"`
	MetricsPort           int      `env:"METRICS_PORT" envDefault:"7000"`
	KafkaBootstrapServers []string `env:"KAFKA_PRODUCER_BOOTSTRAP_SERVERS" envDefault:"localhost:9092"`

	DatabaseConnectUrl string `env:"DATABASE_URL" envDefault:"root:@/test"`
}

func (c *Conf) Register() error {
	err := env.Parse(c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Conf) Validate() error {
	return nil
}

func (c *Conf) Print() interface{} {
	return *c
}
