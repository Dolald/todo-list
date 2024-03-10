package cmd

import (
	"fmt"
	"list/db"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, v := range args {
			newId, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("I don't know such symbols")
			} else {
				ids = append(ids, newId)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Println("Failed to mark")
			} else {
				fmt.Println("Marked", id, "as complited")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
