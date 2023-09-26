package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/api"
	"github.com/tradlwa/xj/cmd/valid"
	"github.com/tradlwa/xj/cmdcontext"
	"github.com/tradlwa/xj/config"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/signal"
	"syscall"
)

func NewCmdAuth(c *cmdcontext.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "auth <domain> <username>",
		Short: "Authenticate for domain",
		Args:  valid.ExpectedArgs(),
		Run: func(cmd *cobra.Command, args []string) {
			password, err := getPassword("Your password: ")
			if err != nil {
				fmt.Println(err)
				return
			}

			domain := args[0]
			username := args[1]

			apiClient := api.NewClient(domain)
			if err = api.Login(apiClient, username, password); err != nil {
				fmt.Println(err)
				return
			}
			c.EnvConfig.Set(c.Env, domain)
			if err = config.WriteConfig(c.EnvConfig); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Authentication successful!")
		},
	}
	return command
}

func getPassword(prompt string) (string, error) {
	initialTermState, err := terminal.GetState(syscall.Stdin)
	if err != nil {
		return "", err
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		_ = terminal.Restore(syscall.Stdin, initialTermState)
		os.Exit(1)
	}()
	fmt.Print(prompt)
	password, err := terminal.ReadPassword(syscall.Stdin)
	fmt.Println("")
	if err != nil {
		return "", err
	}
	signal.Stop(c)
	return string(password), nil
}
