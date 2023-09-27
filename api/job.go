package api

import (
	"github.com/tradlwa/xj/api/urlcodec"
)

type JobData struct {
	ID                     int    `json:"id"`
	JobGroup               int    `json:"jobGroup"`
	JobCron                string `json:"jobCron"`
	JobDesc                string `json:"jobDesc"`
	AddTime                string `json:"addTime"`
	UpdateTime             string `json:"updateTime"`
	Author                 string `json:"author"`
	AlarmEmail             string `json:"alarmEmail"`
	ExecutorRouteStrategy  string `json:"executorRouteStrategy"`
	ExecutorHandler        string `json:"executorHandler"`
	ExecutorParam          string `json:"executorParam"`
	ExecutorBlockStrategy  string `json:"executorBlockStrategy"`
	ExecutorTimeout        int    `json:"executorTimeout"`
	ExecutorFailRetryCount int    `json:"executorFailRetryCount"`
	GlueType               string `json:"glueType"`
	GlueSource             string `json:"glueSource"`
	GlueRemark             string `json:"glueRemark"`
	GlueUpdatetime         string `json:"glueUpdatetime"`
	ChildJobID             string `json:"childJobId"`
	TriggerStatus          int    `json:"triggerStatus"`
	TriggerLastTime        int    `json:"triggerLastTime"`
	TriggerNextTime        int    `json:"triggerNextTime"`
}

func (d JobData) Status() string {
	if d.TriggerStatus == 0 {
		return "off"
	}
	if d.TriggerStatus == 1 {
		return "on"
	}
	return ""
}

type JobOptions struct {
	ID      int    `url:"id"`
	Group   int    `url:"jobGroup"`
	Status  int    `url:"triggerStatus"`
	Desc    string `url:"jobDesc"`
	Handler string `url:"executorHandler"`
	Start   int    `url:"start"`
	Length  int    `url:"length"`
}

func NewJobOptions() *JobOptions {
	return &JobOptions{
		Status: -1,
		Start:  0,
		Length: 10,
	}
}

func JobPage(client *Client, opts *JobOptions) (*PageResponse[JobData], error) {
	var response PageResponse[JobData]
	if err := client.Post("jobinfo/pageList", urlcodec.StructToStringReader(opts), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func JobStart(client *Client, opts *JobOptions) (*BaseResponse, error) {
	var response BaseResponse
	err := client.Post("jobinfo/start", urlcodec.StructToStringReader(opts), &response)
	return &response, err
}

func JobStop(client *Client, opts *JobOptions) (*BaseResponse, error) {
	var response BaseResponse
	err := client.Post("jobinfo/stop", urlcodec.StructToStringReader(opts), &response)
	return &response, err
}

func JobRemove(client *Client, opts *JobOptions) (*BaseResponse, error) {
	var response BaseResponse
	err := client.Post("jobinfo/stop", urlcodec.StructToStringReader(opts), &response)
	return &response, err
}
