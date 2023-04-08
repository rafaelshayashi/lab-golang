package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "golang-fundamentals",
}

func Execute() {
	rootCmd.AddCommand(CollectionsCmd())
	rootCmd.AddCommand(FundamentalsCmd())
	rootCmd.AddCommand(StandardLibCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
