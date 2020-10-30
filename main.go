package main

import (
	"github.com/spf13/cobra"
	"github.com/taisho6339/gobake/cmd"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "gobake",
}

func main() {
	rootCmd.AddCommand(cmd.Up())
	rootCmd.AddCommand(cmd.App())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
