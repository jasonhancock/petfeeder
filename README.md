# petfeeder

A Raspberry Pi powered cat feeder

[Video on YouTube](https://www.youtube.com/watch?v=s813IpMfKfs)

[![backside](https://raw.githubusercontent.com/jasonhancock/petfeeder/master/pics/back_thumb.jpg)](https://raw.githubusercontent.com/jasonhancock/petfeeder/master/pics/back.jpg)

[![profile](https://raw.githubusercontent.com/jasonhancock/petfeeder/master/pics/profile_thumb.jpg)](https://raw.githubusercontent.com/jasonhancock/petfeeder/master/pics/profile.jpg)

## Why?

My family is preparing to take a trip and the folks who usually look after our cat are also out of town. Rather than put him in a kennel, leave out a big bowl of food, etc., I decided to wire up something myself.

## Requirements

* Deliver approximately 1/3 cup of dry cat food three times daily at 06:30, 17:30, and 21:00.

## Hardware

I built this mainly out of components I had on hand. The only thing I needed to purchase was a little bit of lumber, some screws, nylon standoffs, and the cereal dispenser.

* Raspberry Pi Zero
* USB OTG adapter
* USB WiFi Card
* Pin header for Pi Zero
* 5V micro USB power supply
* 12V power supply
* ULN2803 darlington array
* Relay board
* 12V, 6 RPM motor
* Motor mount bracket
* 6mm to 1/4" coupler
* 5" of 1/4" D shaft
* Cereal dispenser
* 1"x3" pine lumber
* Wood screws
* nylon standoffs
* #6 x 1" machine screws
* #6 washers


## Feed Duration

I loaded up the feeder with the target kibble. I ran the feeder for 5 seconds several times and weighed the resulting kibble. On average, the feeder was dispensing 15g of kibble per second. I then weighed out 1/3 cup by volume of kibble. On average it was 42 grams. That means I need to run the machine for approximately 3 seconds at a time, three times a day for my cat to receive the proper amount of food. I'm going to err on the side of caution and run the machine for 3300ms, which should deliver a dose of 49.5 grams of food per feeding.

I also built in a web server that allows you to dispense a feeding on demand. The length of the on-demand feeding is controlled by the `default_duration` parameter in the configuration file. There is no authentication in front of the web server, so run it at your own risk and don't expose it to the internet. I intended to reverse proxy it from my internet facing web server (secured with SSL and basic authentication), but I have yet to complete that portion of the project.

## Golang

I could have easily written this project in Python. I've built several things for Raspberry Pi's that are written in Python. I know enough Python to be dangerous, but I don't work with the language every day. Lately I've been writing a lot of Go...to the point where it's become my primary language for most tasks. As this is a relatively simple/straightforward project, I felt it was a good way to dabble with Go and GPIO on the Pi.

I didn't want to install the Go runtime on the Pi. Thankfully, you can cross compile Go from x86\_64 to arm; this allows me to develop on a x86\_64 Debian machine, package the code, then deploy a built .deb to the Pi. See the [Makefile](Makefile) For more details.
