package cmd

import (
	"fmt"
	"strconv"

	"../consts"
	"../db"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "rm removes a task from the task list",
	Run: func(cmd *cobra.Command, args []string) {
		tasksBucket := consts.TASKS_BUCKET
		ids := []int{}
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Could not understand:", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks(tasksBucket)
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
			err := db.DeleteTask(t.Key, tasksBucket)
			if err != nil {
				fmt.Printf("Task `%d` not removed. Error %s", id, err)
			} else {
				fmt.Printf("Task `%d` is gone.", id)
			}
		}
	},
}
