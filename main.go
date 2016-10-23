package main

import (
	"flag"
	"log"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/stianeikeland/go-rpio"
)

func feed(dur time.Duration) {
	log.Printf("running feed for %s\n", dur)

	time.Sleep(dur)

	log.Println("done running feed")
}

func main() {
	configFile := flag.String("c", "/etc/petfeeder/config.yaml", "configuration file")
	flag.Parse()

	err := rpio.Open()

	if err != nil {
		panic(err)
	}

	pin := rpio.Pin(14)
	pin.Output() // Output mode
	pin.Low()    // Set pin High

	conf, err := loadConfigFile(*configFile)
	if err != nil {
		panic(err)
	}

	for _, v := range conf.Schedule {
		log.Printf("Scheduling a job at %s for %s\n", v.Time, v.Duration)
		gocron.Every(1).Day().At(v.Time).Do(feed, v.Duration)

	}

	<-gocron.Start()
}
