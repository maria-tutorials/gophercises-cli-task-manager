package cmd

import (
	"fmt"
	"strconv"

	"../db"

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

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("ups something went wrong", err)
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task num:", id)
				continue
			}
			t := tasks[id-1]
			err := db.DeleteTask(t.Key)
			if err != nil {
				fmt.Printf("Task `%d` not marked as done. Error %s", id, err)
			} else {
				fmt.Printf("Task `%d` is done.", id)
			}
		}

		fmt.Println(ids)
	},
}
