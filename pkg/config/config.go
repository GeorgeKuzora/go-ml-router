package config

import (
	"fmt"
	"go-ml-router/pkg/fs"
	"net/url"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App App `yaml:"app"`
	Backends []Backend `yaml:"backends"`
}

type App struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Routes Routes `yaml:"routes"`
	readTimeoutInSec int `yaml:"readTimeoutInSec"`
	writeTimeoutInSec int `yaml:"writeTimeoutInSec"`
	idleTimeoutInSec  int `yaml:"idleTimeoutInSec"`
}

func (a App) ReadTimeout() time.Duration {
	return time.Duration(a.readTimeoutInSec) * time.Second
}

func (a App) WriteTimeout() time.Duration {
	return time.Duration(a.writeTimeoutInSec) * time.Second
}

func (a App) IdleTimeout() time.Duration {
	return time.Duration(a.idleTimeoutInSec) * time.Second
}

type Routes struct {
	Predict string `yaml:"predict"`
	Health string `yaml:"health"`
	Metrics string `yaml:"metrics"`
}

type Backend struct {
	Name string `yaml:"name"`
	Priority int `yaml:"priority"`
	Address string `yaml:"address"`
}

func (b Backend) Url() *url.URL {
	u, _ := url.Parse(b.Address)
	return u
}

func FromYaml(path string) (Config, error) {
	data, err := fs.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return Config{}, fmt.Errorf("Failed read config from yaml: %s invalid yaml", path)
	}
	return config, nil
}
