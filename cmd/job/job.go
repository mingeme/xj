package job

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/api"
	"github.com/tradlwa/xj/cmdcontext"
	"os"
)

func NewCmdJob(c *cmdcontext.Context) *cobra.Command {
	taskCmd := &cobra.Command{
		Use:     "job",
		Aliases: []string{"j"},
		Short:   "Management for job",
	}

	taskCmd.AddCommand(NewCmdJobLs(c))
	return taskCmd
}

func NewCmdJobLs(c *cmdcontext.Context) *cobra.Command {
	opts := api.NewJobOptions()
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "search job",
		Run: func(cmd *cobra.Command, args []string) {
			page, err := api.JobPage(c.ApiClient(), opts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				return
			}
			t := table.NewWriter()
			t.AppendHeader(table.Row{"#", "Job Desc", "Executor Handler", "Job Cron", "Author"})
			for _, data := range page.Data {
				t.AppendRow([]interface{}{data.ID, data.JobDesc, data.ExecutorHandler, data.JobCron, data.Author})
			}
			fmt.Println(t.Render())
		},
	}
	lsCmd.Flags().StringVarP(&opts.Handler, "handler", "x", "", "search executor handler")
	lsCmd.Flags().StringVarP(&opts.Desc, "desc", "d", "", "search by job description")
	lsCmd.Flags().IntVarP(&opts.Group, "group", "g", 0, "search by job group")

	return lsCmd
}
