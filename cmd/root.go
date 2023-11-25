package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "targeter",
	Short: "Targeter is a Cli to target resources in a terraform file",
}

func Execute(ctx context.Context) {
	cobra.CheckErr(RootCmd.ExecuteContext(ctx))
}
