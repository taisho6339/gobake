package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taisho6339/gobake/pkg"
	"os"
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

func up() error {
	path, err := pkg.LookupGoPath()
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
}

func Up() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "generate a file to export environments for Golang",
		RunE: func(cmd *cobra.Command, args []string) error {
			return up()
		},
	}
}
