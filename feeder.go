package main

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

type feeder interface {
	feed(duration time.Duration)
}

type Feeder struct {
	pinFeed      rpio.Pin
	pinInterlock rpio.Pin
}

type NullFeeder struct{}

func NewFeeder(feedPinNum, interlockPinNum int) (*Feeder, error) {
	err := rpio.Open()

	if err != nil {
		return nil, err
	}

	f := &Feeder{}

	f.pinFeed = rpio.Pin(feedPinNum)
	f.pinFeed.Output()
	f.pinFeed.Low()

	// Pin 15 is hooked up to relay 2 and wired in normally closed fashion. Thus,
	// when we drive it low here, the relay starts passing the signal. This gates
	// the signal on pin 14 during the boot process, otherwise the output on pin
	// 14 turns the motor on during the boot process for about 11 seconds and the
	// cat gets way over-fed. This protects us against over-feeding in the case
	// of the power being restored after a power-outage.
	f.pinInterlock = rpio.Pin(interlockPinNum)
	f.pinInterlock.Output()
	f.pinInterlock.High()

	return f, nil
}

func (f Feeder) feed(duration time.Duration) {
	log.Println("running feed for ", duration)

	f.pinFeed.High()
	time.Sleep(duration)
	f.pinFeed.Low()

	log.Println("done running feed")
}

func NewNullFeeder() *NullFeeder {
	return &NullFeeder{}
}

func (f NullFeeder) feed(duration time.Duration) {
	log.Println("running feed for ", duration)
	time.Sleep(duration)
	log.Println("done running feed")
}
