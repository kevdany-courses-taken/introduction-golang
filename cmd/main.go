// Package main shows how to use Cobra library in golang.
package main

import (
	"github.com/kevdany-courses-taken/introduction-golang/internal/cli"
	"github.com/spf13/cobra"
)

func main(){
	rootCmd := &cobra.Command{Use: "cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()
}