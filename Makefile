all: test install
	@true

test:
	go test -v

install:
	go install

deps:
	go get github.com/jasonlvhit/gocron
	go get gopkg.in/yaml.v2
