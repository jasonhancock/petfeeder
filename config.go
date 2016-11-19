package main

import (
	"io"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Addr            string        `yaml:"addr"`
	Assets          string        `yaml:"assets"`
	DefaultDuration time.Duration `yaml:"default_duration"`
	Schedule        []Task        `yaml:"schedule"`
	Templates       string        `yaml:"templates"`
}

type Task struct {
	Time     string        `yaml:"time"`
	Duration time.Duration `yaml:"duration"`
}

func loadConfig(r io.Reader) (*Config, error) {
	var c Config

	bytes, err := ioutil.ReadAll(r)
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func loadConfigFile(filename string) (*Config, error) {
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	return loadConfig(fh)
}
