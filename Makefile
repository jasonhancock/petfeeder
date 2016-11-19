all: deps test build
	@true

test:
	go test -v

deps:
	go get github.com/gorilla/mux
	go get github.com/jasonlvhit/gocron
	go get github.com/matryer/temple
	go get github.com/stianeikeland/go-rpio
	go get github.com/tylerb/graceful
	go get gopkg.in/yaml.v2

build:
	GOOS=linux GOARCH=arm go build -o petfeeder
	$(MAKE) -C cmd/manualfeed

package:
	$(MAKE) -C debian

up: down
	go build
	./petfeeder -dev &

down:
	killall petfeeder || true
