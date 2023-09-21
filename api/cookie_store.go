package api

import (
	"encoding/json"
	"fmt"
	"github.com/tradlwa/xj/config"
	"net/http"
)

type CookieStore interface {
	Save(host string, cookie http.Cookie)
	Get(host string) ([]*http.Cookie, error)
}

type cookieStore struct {
}

func NewCookieStore() CookieStore {
	return &cookieStore{}
}

func (c *cookieStore) Save(host string, cookie http.Cookie) {
	//TODO implement me
	panic("implement me")
}

func (c *cookieStore) Get(host string) ([]*http.Cookie, error) {
	bytes, err := config.ReadCookieFile()
	if err != nil {
		return nil, err
	}
	hostCookie := HostCookie{}
	if err := json.Unmarshal(bytes, &hostCookie); err != nil {
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

type HostCookie map[string][]CookieItem

type CookieItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
