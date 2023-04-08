package cmd

import (
	"github.com/rafaelshayashi/lab-golang/fundamentals/basics"
	"github.com/spf13/cobra"
)

func FundamentalsCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "basics",
		Run: func(cmd *cobra.Command, args []string) {
			basics.WorkingWithPrimitives()
			basics.WorkingWithPointer()
			basics.WorkingWithStructs()
		},
	}
	return &cmd
}
