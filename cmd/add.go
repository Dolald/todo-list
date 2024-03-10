package cmd

import (
	"fmt"
	"list/db"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		newTask := strings.Join(args, " ")
		_, err := db.CreateTask(newTask)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your list\n", newTask)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
