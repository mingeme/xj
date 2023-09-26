package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tradlwa/xj/config"
	"net/http"
	"os"
)

type CookieStore interface {
	Save(host string, cookie *http.Cookie) error
	Get(host string) ([]*http.Cookie, error)
}

type cookieStore struct {
}

func NewCookieStore() CookieStore {
	return &cookieStore{}
}

func (c *cookieStore) Save(host string, cookie *http.Cookie) error {
	hostCookie, err := readFile()
	if err != nil && !errors.Is(err, config.ErrCookieFileNotFound) {
		return err
	}
	items := []CookieItem{{
		Name:  cookie.Name,
		Value: cookie.Value,
	}}
	if errors.Is(err, config.ErrCookieFileNotFound) {
		hostCookie = make(HostCookie)
	}
	hostCookie[host] = items
	b, err := json.Marshal(hostCookie)
	if err != nil {
		return err
	}
	return os.WriteFile(config.CookieFilePath(), b, os.ModePerm)
}

func (c *cookieStore) Get(host string) ([]*http.Cookie, error) {
	hostCookie, err := readFile()
	if err != nil {
		return nil, err
	}
	if items, ok := hostCookie[host]; ok {
		var cookies []*http.Cookie
		for _, item := range items {
			cookies = append(cookies, &http.Cookie{Name: item.Name, Value: item.Value})
		}
		return cookies, nil
	}

	return nil, fmt.Errorf("no cookie set for host %s", host)
}

func readFile() (HostCookie, error) {
	bytes, err := config.ReadCookieFile()
	if err != nil {
		return nil, err
	}
	hostCookie := HostCookie{}
	if err := json.Unmarshal(bytes, &hostCookie); err != nil {
		return nil, err
	}
	return hostCookie, nil
}

type HostCookie map[string][]CookieItem

type CookieItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
