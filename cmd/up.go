package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/template"
)

const (
	envFileName = ".gobake"
	envTemplate = `
export GOPATH={{.GoPath}}
export PATH=$PATH:$GOPATH/bin
export GOENV_DISABLE_GOPATH=1
`
)

type templateArgs struct {
	GoPath string
}

func lookupGoPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dirs := strings.Split(wd, "/src/")
	if len(dirs) < 2 {
		return "", errors.New("current path is invalid")
	}
	return dirs[0], nil
}

func Up() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "generate a file to export environments for Golang",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := lookupGoPath()
			if err != nil {
				return err
			}
			file, err := os.Create(envFileName)
			if err != nil {
				return err
			}
			tl, err := template.New(envFileName).Parse(envTemplate)
			if err != nil {
				return err
			}
			if err := tl.Execute(file, templateArgs{GoPath: path}); err != nil {
				return err
			}
			return nil
		},
	}
}
