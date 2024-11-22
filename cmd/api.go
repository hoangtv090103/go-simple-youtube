package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func apiCmd() *cobra.Command {
	return &cobra.Command{
		Use: "api",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, world!")
		},
	}
}
