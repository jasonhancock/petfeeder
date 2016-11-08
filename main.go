package main

import (
	"flag"
	"log"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/stianeikeland/go-rpio"
)

func feed(dur time.Duration, pin rpio.Pin) {
	log.Printf("running feed for %s\n", dur)

	pin.High()
	time.Sleep(dur)
	pin.Low()

	log.Println("done running feed")
}

func main() {
	configFile := flag.String("c", "/etc/petfeeder/config.yaml", "configuration file")
	flag.Parse()

	err := rpio.Open()

	if err != nil {
		panic(err)
	}

	// Pin 15 is hooked up to relay 2 and wired in normally closed fashion. Thus,
	// when we drive it low here, the relay starts passing the signal. This gates
	// the signal on pin 14 during the boot process, otherwise the output on pin
	// 14 turns the motor on during the boot process for about 11 seconds and the
	// cat gets way over-fed. This protects us against over-feeding in the case
	// of the power being restored after a power-outage.
	interlockPin := rpio.Pin(15)
	interlockPin.Output()
	interlockPin.Low()

	pin := rpio.Pin(14)
	pin.Output()
	pin.Low()

	conf, err := loadConfigFile(*configFile)
	if err != nil {
		panic(err)
	}

	for _, v := range conf.Schedule {
		log.Printf("Scheduling a job at %s for %s\n", v.Time, v.Duration)
		gocron.Every(1).Day().At(v.Time).Do(feed, v.Duration, pin)
	}

	<-gocron.Start()
}
