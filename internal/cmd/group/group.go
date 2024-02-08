package group

import (
	"fmt"
	"github.com/heminghu/xj/internal/api"
	"github.com/heminghu/xj/internal/cmdcontext"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
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
	opts := api.NewGroupOptions()
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "search job group",
		Run: func(cmd *cobra.Command, args []string) {
			page, err := api.GroupPage(c.ApiClient(), opts)
			if err != nil {
				fmt.Printf("%+v", err)
				return
			}
			t := table.NewWriter()
			t.AppendHeader(table.Row{"#", "App Name", "Title"})
			for _, data := range page.Data {
				t.AppendRow([]any{data.ID, data.App, data.Title})
			}
			fmt.Println(t.Render())
		},
	}
	cmd.Flags().StringVarP(&opts.App, "app", "a", "", "app name")
	cmd.Flags().StringVarP(&opts.Title, "title", "t", "", "title")
	cmd.Flags().IntVarP(&opts.Start, "start", "s", 0, "page start")
	cmd.Flags().IntVarP(&opts.Length, "len", "l", 10, "page length")

	return cmd
}
