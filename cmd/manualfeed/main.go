package main

import (
	"flag"
	"log"
	"time"

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
	dur := flag.Duration("d", 0, "duration to run feeder")
	flag.Parse()

	err := rpio.Open()

	if err != nil {
		panic(err)
	}

	pin := rpio.Pin(14)
	pin.Output() // Output mode
	pin.Low()    // Set pin Low

	feed(*dur, pin)
}
