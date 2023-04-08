package cmd

import (
	"github.com/rafaelshayashi/lab-golang/fundamentals/collections"
	"github.com/spf13/cobra"
)

func CollectionsCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "collections",
		Run: func(cmd *cobra.Command, args []string) {
			collections.WorkingWithArray()
			collections.WorkingWithSlice()
			collections.IteratingOverArrayOfInt()
			collections.FilterArray()
		},
	}
	return &cmd
}
