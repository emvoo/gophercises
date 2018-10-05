package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"gophercises/07/task/store"
	"fmt"
	"strings"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Args:  cobra.MaximumNArgs(100),
	Run: func(cmd *cobra.Command, args []string) {
		s := store.OpenDB()

		arg := strings.Join(args, " ")

		if err := s.InsertTask(arg); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Added \"%s\" to your task list.\n", arg)
	},
}
