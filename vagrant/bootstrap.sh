#!/bin/bash

apt-get install -y git

# Install Go
goversion=1.7.3
gofile=go${goversion}.linux-amd64.tar.gz
gourl=https://storage.googleapis.com/golang/${gofile}
wget -q -O /usr/local/${gofile} ${gourl}
mkdir /usr/local/go
tar -xzf /usr/local/${gofile} -C /usr/local/go --strip 1


# Set up go environment
echo 'export PATH=$PATH:/usr/local/go/bin' > /etc/profile.d/golang.sh
echo 'export GOPATH=/home/vagrant/go' >> /etc/profile.d/golang.sh
echo 'export petfeeder=$GOPATH/src/github.com/jasonhancock/petfeeder' >> /etc/profile.d/golang.sh
. /etc/profile.d/golang.sh

# Link the mf2 directory into the $GOPATH
mkdir -p $GOPATH/src/github.com/jasonhancock
chown -R vagrant.vagrant $GOPATH
ln -s /vagrant $GOPATH/src/github.com/jasonhancock/petfeeder

# Create the config path, link in the example config:
mkdir /etc/petfeeder
ln -s $GOPATH/src/github.com/jasonhancock/petfeeder/config.yaml /etc/petfeeder/config.yaml
