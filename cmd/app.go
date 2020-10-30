package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	defaultRegistry = "github.com"
)

func app(registry string, args []string) error {
	authorName := args[0]
	appName := args[1]
	dirName := fmt.Sprintf("%s/src/%s/%s/%s", appName, registry, authorName, appName)
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		return err
	}
	if err := os.Chdir(dirName); err != nil {
		return err
	}
	if err := up(); err != nil {
		return err
	}
	return nil
}

func App() *cobra.Command {
	var (
		registry string
	)
	command := &cobra.Command{
		Use:   "app [author name] [app name] <flags>",
		Short: "create the directory structure for a new application",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app(registry, args)
		},
	}
	command.Flags().StringVar(&registry, "registry", defaultRegistry, "registry name which registers your application. eg. github.com")
	return command
}
