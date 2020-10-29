package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/taisho6339/gobake/cmd"
)

var rootCmd = &cobra.Command{
	Use: "gobake",
}

func main() {
	rootCmd.AddCommand(cmd.Up())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
