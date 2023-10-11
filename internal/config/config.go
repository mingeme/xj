package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var ErrCookieFileNotFound = fmt.Errorf("cookie file not found")

func Dir() string {
	d, _ := os.UserHomeDir()
	return filepath.Join(d, ".config", "xj")
}

func ReadCookieFile() ([]byte, error) {
	filename := CookieFilePath()
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, ErrCookieFileNotFound
	}
	if err != nil {
		return nil, err
	}
	return os.ReadFile(filename)
}

func CookieFilePath() string {
	return filepath.Join(Dir(), "cookies.json")
}
