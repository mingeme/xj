package api

import (
	"github.com/tradlwa/xj/internal/api/urlcodec"
)

type GroupData struct {
	ID           int      `json:"id"`
	App          string   `json:"appname"`
	Title        string   `json:"title"`
	AddressType  int      `json:"addressType"`
	AddressList  string   `json:"addressList"`
	RegistryList []string `json:"registryList"`
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

func GroupPage(client *Client, opts *GroupOptions) (*PageResponse[GroupData], error) {
	var response PageResponse[GroupData]
	if err := client.Post("jobgroup/pageList", urlcodec.StructToStringReader(opts), &response); err != nil {
		return nil, err
	}
	return &response, nil
}
