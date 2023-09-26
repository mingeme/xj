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
	cmd := &cobra.Command{
		Use:     "job",
		Aliases: []string{"j"},
		Short:   "Management for job",
	}

	cmd.AddCommand(NewCmdJobLs(c))
	cmd.AddCommand(NewCmdTrigger(c))
	return cmd
}

func NewCmdJobLs(c *cmdcontext.Context) *cobra.Command {
	opts := api.NewJobOptions()
	cmd := &cobra.Command{
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
	cmd.Flags().StringVarP(&opts.Handler, "handler", "x", "", "search executor handler")
	cmd.Flags().StringVarP(&opts.Desc, "desc", "d", "", "search by job description")
	cmd.Flags().IntVarP(&opts.Group, "group", "g", 0, "search by job group")

	return cmd
}

func NewCmdTrigger(c *cmdcontext.Context) *cobra.Command {
	opts := api.NewTriggerOptions()
	cmd := &cobra.Command{
		Use:     "trigger",
		Aliases: []string{"t"},
		Short:   "trigger job",
		Run: func(cmd *cobra.Command, args []string) {
			response, err := api.TriggerJob(c.ApiClient(), opts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				return
			}
			if response.Code == 200 {
				fmt.Println("triggered")
			} else {
				fmt.Printf("%v", response)
			}
		},
	}
	cmd.Flags().IntVarP(&opts.ID, "id", "i", -1, "job id")
	cmd.Flags().StringVar(&opts.Param, "param", "p", "job parameter")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
