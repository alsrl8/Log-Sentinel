package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
	"sync"
)

type Config struct {
	Watch []struct {
		Path string `yaml:"path"`
		Name string `yaml:"name"`
	} `yaml:"watch"`
}

var (
	instance *Config
	once     sync.Once
)

func LoadConfig() (*Config, error) {
	var err error
	once.Do(func() {
		configFilePath := getLogSentinelConfigPath()
		if configFilePath == "" {
			err = errors.New("config file path is empty")
			return
		}
		log.Println("Loading config from", configFilePath)
		instance, err = loadConfig(configFilePath)
		printWatchList(instance)
	})
	return instance, err
}

func getLogSentinelConfigPath() string {
	return os.Getenv("LOG_SENTINEL_CONFIG_PATH")
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

func printWatchList(cfg *Config) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Watching %d files\n", len(cfg.Watch)))
	for _, w := range cfg.Watch {
		sb.WriteString(fmt.Sprintf("\t%s\t%s\n", w.Name, w.Path))
	}

	log.Println(sb.String())
}
