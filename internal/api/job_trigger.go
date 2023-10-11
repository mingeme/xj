package api

import (
	"github.com/tradlwa/xj/internal/api/urlcodec"
)

type TriggerOptions struct {
	ID    int    `url:"id"`
	Param string `url:"executorParam"`
}

func NewTriggerOptions() *TriggerOptions {
	return &TriggerOptions{}
}

func TriggerJob(client *Client, opts *TriggerOptions) (*BaseResponse, error) {
	var response BaseResponse
	if err := client.Post("jobinfo/trigger", urlcodec.StructToStringReader(opts), &response); err != nil {
		return nil, err
	}
	return &response, nil
}
