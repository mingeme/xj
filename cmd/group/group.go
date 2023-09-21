package group

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/api"
	"github.com/tradlwa/xj/cmdutil"
)

func NewCmdGroup(c *cmdutil.Context) *cobra.Command {
	groupCmd := &cobra.Command{
		Use:     "group",
		Aliases: []string{"g"},
		Short:   "Management for job group",
	}

	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "执行器查询",
		Run: func(cmd *cobra.Command, args []string) {
			page := api.GroupPage(c.ApiClient, &api.GroupOptions{Start: 0, Length: 10})
			t := table.NewWriter()
			t.AppendHeader(table.Row{"#", "App Name", "Title"})
			for _, data := range page.Data {
				t.AppendRow([]interface{}{data.ID, data.AppName, data.Title})
			}
			fmt.Println(t.Render())
		},
	}

	groupCmd.AddCommand(lsCmd)
	return groupCmd
}
