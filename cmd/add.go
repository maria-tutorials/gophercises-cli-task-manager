package cmd

import (
	"fmt"
	"log"
	"strings"

	"../consts"
	"../db"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task, consts.TASKS_BUCKET)
		if err != nil {
			log.Fatal("ups something went wrong", err)
		}
		fmt.Printf("Added `%s` to the list. \n", task)
	},
}
