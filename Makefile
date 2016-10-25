all: deps test build
	@true

test:
	go test -v

deps:
	go get github.com/jasonlvhit/gocron
	go get gopkg.in/yaml.v2
	go get github.com/stianeikeland/go-rpio

build:
	GOOS=linux GOARCH=arm go build -o petfeeder
	$(MAKE) -C cmd/manualfeed

package:
	$(MAKE) -C debian

