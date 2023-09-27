package group

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/api"
	"github.com/tradlwa/xj/cmdcontext"
)

func NewCmdGroup(c *cmdcontext.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "group",
		Aliases: []string{"g"},
		Short:   "Management for job group",
	}

	cmd.AddCommand(NewCmdGroupLs(c))
	return cmd
}

func NewCmdGroupLs(c *cmdcontext.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "search job group",
		Run: func(cmd *cobra.Command, args []string) {
			page, err := api.GroupPage(c.ApiClient(), api.NewGroupOptions())
			if err != nil {
				fmt.Printf("%+v", err)
				return
			}
			t := table.NewWriter()
			t.AppendHeader(table.Row{"#", "App Name", "Title"})
			for _, data := range page.Data {
				t.AppendRow([]any{data.ID, data.AppName, data.Title})
			}
			fmt.Println(t.Render())
		},
	}
}
