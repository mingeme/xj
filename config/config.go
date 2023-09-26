package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var ErrCookieFileNotFound = fmt.Errorf("cookie file not found")

func Dir() string {
	d, _ := os.UserHomeDir()
	//
	//// Check if directory already exists
	//_, err := os.Stat(path)
	//if err == nil {
	//	return path
	//}
	//
	//// Create the directory
	//err = os.MkdirAll(path, 0755)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Failed to create config directory: %v", err)
	//	return path
	//}

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
