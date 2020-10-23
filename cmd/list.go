package cmd

import (
	"fmt"
	"log"

	"../db"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list prints all the tasks that need to be done",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			log.Fatal("ups something went wrong", err)
		}

		if len(tasks) == 0 {
			fmt.Println("There's no tasks, you're all free")
			return
		}

		fmt.Println("You have the following tasks:")
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", (i + 1), t.Value)
		}

	},
}
