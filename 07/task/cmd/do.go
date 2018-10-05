package cmd

import (
	"github.com/spf13/cobra"
	"gophercises/07/task/store"
		"log"
	"fmt"
	)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("Task number should be passed as an argument.")
			return
		}
		arg := args[0]
		s := store.OpenDB()

		key := []byte(fmt.Sprintf("task%s", arg))
		value, err := s.GetTask(key)
		if err != nil {
			log.Fatal(err)
		}

		if err = s.DeleteTask(key); err != nil {
			log.Fatal(err)
		}

		if err = s.MarkDone(key, value); err != nil {
			log.Fatal(err)
		}

		log.Printf("Task (%s) has been marked done.\n", arg)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
