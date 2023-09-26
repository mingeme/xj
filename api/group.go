package api

import (
	"github.com/tradlwa/xj/api/urlcodec"
	"strings"
)

type GroupResponse struct {
	RecordsFiltered int `json:"recordsFiltered"`
	Data            []struct {
		ID           int      `json:"id"`
		AppName      string   `json:"appname"`
		Title        string   `json:"title"`
		AddressType  int      `json:"addressType"`
		AddressList  string   `json:"addressList"`
		RegistryList []string `json:"registryList"`
	} `json:"data"`
	RecordsTotal int `json:"recordsTotal"`
}

type GroupOptions struct {
	App    string `url:"appname"`
	Title  string `url:"title"`
	Start  int    `url:"start"`
	Length int    `url:"length"`
}

func NewGroupOptions() *GroupOptions {
	return &GroupOptions{
		Start:  0,
		Length: 10,
	}
}

func GroupPage(client *Client, opts *GroupOptions) (*GroupResponse, error) {
	values := urlcodec.StructToValues(opts)
	var response GroupResponse
	if err := client.Post("jobgroup/pageList", strings.NewReader(values.Encode()), &response); err != nil {
		return nil, err
	}
	return &response, nil
}
