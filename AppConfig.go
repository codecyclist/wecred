package main

import (
	"errors"
	"flag"
)

type AppConfig struct {
	Url string
}

func BuildAppConfig() (config AppConfig, err error) {
	config = AppConfig{}
	err = nil

	flag.StringVar(&config.Url, "url", "", "The URL of your site")
	flag.Parse()
	if config.Url == "" {
		flag.Usage()
		err = errors.New("Please provide a valid site URI.")
	}

	return config, err
}
