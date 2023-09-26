package api

import (
	"github.com/tradlwa/xj/api/urlcodec"
	"strings"
)

type TriggerOptions struct {
	ID    int    `url:"id"`
	Param string `url:"executorParam"`
}

func NewTriggerOptions() *TriggerOptions {
	return &TriggerOptions{}
}

func TriggerJob(client *Client, opts *TriggerOptions) (*BaseResponse, error) {
	values := urlcodec.StructToValues(opts)
	var response BaseResponse
	if err := client.Post("jobinfo/trigger", strings.NewReader(values.Encode()), &response); err != nil {
		return nil, err
	}
	return &response, nil
}
