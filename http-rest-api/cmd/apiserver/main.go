package main

import (
	"flag"
	apiserver2 "github.com/HenRok1/http-rest-api/internal/app/apiserver"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver2.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver2.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
