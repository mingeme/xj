package cmdcontext

import (
	"fmt"
	"os"
	"strconv"

	"github.com/heminghu/xj/internal/api"
	"github.com/heminghu/xj/internal/config"
)

type Context struct {
	EnvConfig *config.EnvConfig

	Env string
}

func (c *Context) Domain() string {
	domain, err := c.EnvConfig.Get(c.Env)
	if err != nil {
		fmt.Printf("failed to read env: %v", err)
		os.Exit(1)
	}
	return domain
}

func (c *Context) JobInfo() string {
	return c.Domain() + "xxl-job-admin/jobinfo"
}

func (c *Context) JobLog(id int) string {
	return c.Domain() + "xxl-job-admin/joblog?jobId=" + strconv.Itoa(id)
}

func (c *Context) ApiClient() *api.Client {
	return api.NewClient(c.Domain())
}
