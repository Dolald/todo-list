package cmd

import (
	"fmt"
	"list/db"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks, you have taking a vacation")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Println(i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
