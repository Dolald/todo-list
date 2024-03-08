package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		ids := make([]int, 0)
		for _, v := range args {
			newId, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("I don't know such symbols")
			} else {
				ids = append(ids, newId)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
