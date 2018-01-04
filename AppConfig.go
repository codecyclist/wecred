package main

import (
	"errors"
	"flag"
)

type AppConfig struct {
	Url     string
	Verbose bool
}

func BuildAppConfig() (config AppConfig, err error) {
	flag.BoolVar(&config.Verbose, "verbose", false, "Provide verbose output")

	if flag.Parse(); len(flag.Args()) == 1 {
		config.Url = flag.Args()[0]
	} else {
		flag.Usage()
		err = errors.New("Please provide a valid site URL (i.E. http://example.com")
	}
	return
}
