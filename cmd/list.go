package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list prints all the tasks that need to be done",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("here they are")
	},
}
