package main

import (
	"flag"
	"log"
)

func main() {
	configFile := flag.String("c", "/etc/petfeeder/config.yaml", "configuration file")
	isDev := flag.Bool("dev", false, "is development?")
	flag.Parse()

	var fdr feeder
	var err error
	if *isDev {
		fdr = NewNullFeeder()
	} else {
		fdr, err = NewFeeder(14, 15)
		if err != nil {
			log.Fatal(err)
		}
	}

	conf, err := loadConfigFile(*configFile)
	if err != nil {
		panic(err)
	}

	server, err := NewServer(conf, fdr)

	if err != nil {
		log.Fatal(err)
	}

	server.Run(conf.Addr)
}
