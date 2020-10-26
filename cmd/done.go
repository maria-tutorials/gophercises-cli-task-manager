package cmd

import (
	"fmt"
	"log"

	"../consts"
	"../db"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(doneCmd)
}

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "done prints all the tasks that have been completed",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks(consts.COMPLETED_BUCKET)
		if err != nil {
			log.Fatal("ups something went wrong", err)
		}

		if len(tasks) == 0 {
			fmt.Println("You have not finished a single task, better get to it")
			return
		}

		fmt.Println("You have completed the following tasks:")
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", (i + 1), t.Value)
		}

	},
}
