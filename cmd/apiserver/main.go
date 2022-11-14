package main

import (
	"flag"
	"http-rest/cmd/interval/app/apiserver"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/pelletier/go-toml/v2"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to config file")
}

func main() {
	flag.Parsed()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
