package config

import (
	"fmt"
	"go-ml-router/pkg/fs"
	"net/url"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App App `yaml:"app"`
	Backends []Backend `yaml:"backends"`
}

type App struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
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
