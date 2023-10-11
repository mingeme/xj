package cmdcontext

import (
	"fmt"
	"os"

	"github.com/tradlwa/xj/internal/api"
	"github.com/tradlwa/xj/internal/config"
)

type Context struct {
	EnvConfig *config.EnvConfig

	Env string
}

func (c *Context) ApiClient() *api.Client {
	domain, err := c.EnvConfig.Get(c.Env)
	if err != nil {
		fmt.Printf("failed to read env: %v", err)
		os.Exit(1)
	}
	return api.NewClient(domain)
}
