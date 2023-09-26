package api

import (
	"github.com/tradlwa/xj/api/urlcodec"
	"strings"
)

type JobOptions struct {
	App    string `url:"appname"`
	Title  string `url:"title"`
	Start  int    `url:"start"`
	Length int    `url:"length"`
}

func JobPage(client *Client, opts *JobOptions) {
	values := urlcodec.StructToValues(opts)
	var response any
	client.Post("jobgroup/pageList", strings.NewReader(values.Encode()), &response)
}
