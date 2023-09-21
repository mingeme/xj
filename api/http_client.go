package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	client      *http.Client
	domain      string
	cookieStore CookieStore
}

func NewClient(domain string) *Client {
	return &Client{domain: domain, client: http.DefaultClient, cookieStore: NewCookieStore()}
}

func (c *Client) Do(method string, path string, body io.Reader, response any) error {
	url := restUrl(c.domain, path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	cookies, err := c.cookieStore.Get(req.Host)
	if err != nil {
		return err
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	success := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !success {
		defer resp.Body.Close()
		return fmt.Errorf("failed request %+v", resp)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Post(path string, body io.Reader, response any) error {
	return c.Do(http.MethodPost, path, body, response)
}

func restUrl(domain string, path string) string {
	return fmt.Sprintf("%s/xxl-job-admin/%s", strings.TrimSuffix(domain, "/"), path)
}
