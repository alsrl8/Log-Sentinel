package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Config struct {
	Kube []struct {
		Name         string `yaml:"name"`
		NameSpace    string `yaml:"namespace"`
		ServiceLabel string `yaml:"serviceLabel"`
		LogPath      string `yaml:"logPath"`
		Format       string `yaml:"format"`
	}
}

var (
	instance *Config
	once     sync.Once
)

func LoadConfig() (*Config, error) {
	var e error
	once.Do(func() {
		configFilePath, err := getLogSentinelConfigPath()
		if err != nil {
			e = err
			return
		}
		instance, err = loadConfig(configFilePath)
		if err != nil {
			e = err
			return
		}
	})
	return instance, e
}

func getLogSentinelConfigPath() (string, error) {
	ret := os.Getenv("LOG_SENTINEL_CONFIG_PATH")
	if ret == "" {
		return "", errors.New("LOG_SENTINEL_CONFIG_PATH is not set")
	}
	return ret, nil
}

func loadConfig(configFilePath string) (*Config, error) {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
