package job

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/internal/api"
	"github.com/tradlwa/xj/internal/cmd/validator"
	"github.com/tradlwa/xj/internal/cmdcontext"
)

func NewCmdJob(c *cmdcontext.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "job",
		Aliases: []string{"j"},
		Short:   "Management for job",
	}

	cmd.AddCommand(NewCmdJobLs(c))
	cmd.AddCommand(NewCmdJobTrigger(c))
	cmd.AddCommand(NewCmdJobStart(c))
	cmd.AddCommand(NewCmdJobStop(c))
	cmd.AddCommand(NewCmdJobRemove(c))
	return cmd
}

func NewCmdJobLs(c *cmdcontext.Context) *cobra.Command {
	opts := api.NewJobOptions()
	var nav bool
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "search job",
		Run: func(cmd *cobra.Command, args []string) {
			if nav {
				fmt.Println(c.JobInfo())
			}
			page, err := api.JobPage(c.ApiClient(), opts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				return
			}
			t := table.NewWriter()
			t.AppendHeader(table.Row{"#", "Job Desc", "Executor Handler", "Job Cron", "Author", "Status"})
			for _, data := range page.Data {
				t.AppendRow([]any{data.ID, data.JobDesc, data.ExecutorHandler, data.JobCron, data.Author, data.Status()})
			}
			fmt.Println(t.Render())
		},
	}

	cmd.Flags().StringVarP(&opts.Handler, "handler", "x", "", "executor handler")
	cmd.Flags().StringVarP(&opts.Desc, "desc", "d", "", "job description")
	cmd.Flags().IntVarP(&opts.Group, "group", "g", 0, "job group")
	cmd.Flags().IntVarP(&opts.Status, "status", "", 0, "job status (o/off 1/on)")
	cmd.Flags().IntVarP(&opts.Start, "start", "s", 0, "page start")
	cmd.Flags().IntVarP(&opts.Length, "len", "l", 10, "page length")
	cmd.Flags().BoolVar(&nav, "nav", false, "show navigation")

	return cmd
}

func NewCmdJobTrigger(c *cmdcontext.Context) *cobra.Command {
	opts := api.NewTriggerOptions()
	var nav bool
	cmd := &cobra.Command{
		Use:     "trigger {<job-id>}",
		Aliases: []string{"t"},
		Short:   "trigger job",
		Args:    validator.ExpectedArgs(),
		Run: func(cmd *cobra.Command, args []string) {
			if !parseJobId(&opts.ID, args[0]) {
				return
			}
			if nav {
				fmt.Println(c.JobLog(opts.ID))
			}

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
	cmd.Flags().StringVarP(&opts.Param, "param", "p", "", "job parameter")
	cmd.Flags().BoolVar(&nav, "nav", false, "show navigation")
	return cmd
}

func NewCmdJobStart(c *cmdcontext.Context) *cobra.Command {
	opts := api.NewJobOptions()
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start job {<job-id>}",
		Args:  validator.ExpectedArgs(),
		Run: func(cmd *cobra.Command, args []string) {
			if !parseJobId(&opts.ID, args[0]) {
				return
			}

			response, err := api.JobStart(c.ApiClient(), opts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				return
			}
			if response.Code == 200 {
				fmt.Println("it's started")
			} else {
				fmt.Printf("%v", response)
			}
		},
	}
	return cmd
}

func NewCmdJobStop(c *cmdcontext.Context) *cobra.Command {
	opts := api.NewJobOptions()
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "stop job {<job-id>}",
		Args:  validator.ExpectedArgs(),
		Run: func(cmd *cobra.Command, args []string) {
			if !parseJobId(&opts.ID, args[0]) {
				return
			}

			response, err := api.JobStop(c.ApiClient(), opts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				return
			}
			if response.Code == 200 {
				fmt.Println("it's stopped")
			} else {
				fmt.Printf("%v", response)
			}
		},
	}
	return cmd
}

func NewCmdJobRemove(c *cmdcontext.Context) *cobra.Command {
	opts := api.NewJobOptions()
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "remove job {<job-id>}",
		Args:  validator.ExpectedArgs(),
		Run: func(cmd *cobra.Command, args []string) {
			if !parseJobId(&opts.ID, args[0]) {
				return
			}

			response, err := api.JobRemove(c.ApiClient(), opts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				return
			}
			if response.Code == 200 {
				fmt.Println("it's removed")
			} else {
				fmt.Printf("%v", response)
			}
		},
	}
	return cmd
}

func parseJobId(id *int, arg string) bool {
	val, err := validator.RequireInt(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "required argument <job-id> '%s' should be int", arg)
		return false
	}
	*id = val
	return true
}
