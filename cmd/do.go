package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Long:  "Give the task ids as params",
	Run: func(cmd *cobra.Command, args []string) {
		ids := []int{}
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Could not understand:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}
