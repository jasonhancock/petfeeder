all: test install
	@true

test:
	go test -v

install:
	go install

deps:
	go get github.com/jasonlvhit/gocron
	go get gopkg.in/yaml.v2
	go get github.com/stianeikeland/go-rpio

build:
	GOOS=linux GOARCH=arm go build
