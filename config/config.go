package config

import (
	"fmt"
	"os"
)

var ErrCookieFileNotFound = fmt.Errorf("cookie file not found")

type Config struct {
	Domain string `yaml:"domain"`
}

func Dir() string {
	dirPath := os.Getenv("HOME") + "/.config/xj"

	// Check if directory already exists
	_, err := os.Stat(dirPath)
	if err == nil {
		return dirPath
	}

	// Create the directory
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create config directory: %v", err)
		return dirPath
	}

	return dirPath
}

func ReadCookieFile() ([]byte, error) {
	filename := Dir() + "/cookies.json"

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, ErrCookieFileNotFound
	}
	if err != nil {
		return nil, err
	}
	return os.ReadFile(filename)
}
