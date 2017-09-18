package nsqmate

import "github.com/nsqio/go-nsq"

type Config struct {
	*nsq.Config
}

func NewConfig() *Config {
	return &Config{nsq.NewConfig()}
}
