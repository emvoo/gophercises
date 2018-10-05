package cmd

import (
	"github.com/spf13/cobra"
	"gophercises/07/task/store"
		"log"
	"fmt"
	)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		s := store.OpenDB()

		tasks, err := s.Load()
		if err != nil {
			log.Fatal(err)
		}

		for k, t := range tasks {
			fmt.Printf("%d. %s\n", k+1, t)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
