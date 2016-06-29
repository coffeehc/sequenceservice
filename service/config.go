package main

import "github.com/coffeehc/microserviceboot/consultool"

type Config struct {
	ConsulConfig *consultool.ConsulConfig `yaml:"consul"`
}
