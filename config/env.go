package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
)

var ErrKeyNotFound = errors.New("key not found")

type EnvConfig struct {
	envMap map[string]string
}

func (c *EnvConfig) Get(key string) (string, error) {
	if domain, ok := c.envMap[key]; ok {
		return domain, nil
	}
	return "", ErrKeyNotFound
}

func (c *EnvConfig) Set(key string, domain string) {
	c.envMap[key] = domain
}

func (c *EnvConfig) Serialize() []byte {
	var lines []string
	for key, domain := range c.envMap {
		lines = append(lines, fmt.Sprintf("%s=%s", key, domain))
	}
	return []byte(strings.Join(lines, "\n"))
}

func EnvFile() string {
	return filepath.Join(Dir(), "env")
}

func ReadConfig() (*EnvConfig, error) {
	_, err := os.Stat(EnvFile())
	if err != nil && os.IsNotExist(err) {
		return &EnvConfig{envMap: make(map[string]string)}, nil
	}
	if err != nil {
		return nil, err
	}
	envMap, err := godotenv.Read(EnvFile())
	if err != nil {
		return nil, err
	}
	return &EnvConfig{
		envMap,
	}, nil
}

func WriteConfig(c *EnvConfig) error {
	if err := os.MkdirAll(Dir(), 0755); err != nil {
		return err
	}
	return os.WriteFile(EnvFile(), c.Serialize(), os.FileMode(0644))
}
