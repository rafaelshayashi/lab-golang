package cmd

import (
	"github.com/rafaelshayashi/lab-golang/fundamentals/std_library"
	"github.com/spf13/cobra"
)

func StandardLibCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "standard-lib",
		Run: func(cmd *cobra.Command, args []string) {
			std_library.WorkingWithPrintf()
			std_library.WorkingWithEncoding()
		},
	}
	return &cmd
}
