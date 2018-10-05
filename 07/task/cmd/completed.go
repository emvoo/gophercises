package cmd

import (
	"github.com/spf13/cobra"
	"gophercises/07/task/store"
	"log"
	"fmt"
)

var completedCmd = &cobra.Command{
	Use:"completed",
	Short:"Displays completed tasks.",
	Args:cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		s := store.OpenDB()
		tasks, err := s.LoadCompletedTasks()
		if err != nil {
			log.Fatal(err)
		}

		for k, t := range tasks {
			fmt.Printf("%d. %s\n", k+1, t.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
