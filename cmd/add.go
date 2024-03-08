package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		newTask := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your list\n", newTask)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
