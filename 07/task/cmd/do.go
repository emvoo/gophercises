package cmd

import (
	"github.com/spf13/cobra"
	"gophercises/07/task/store"
		"log"
	"fmt"
	"strconv"
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
		arg, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		s := store.OpenDB()

		tasks, err := s.LoadToDos()
		if err != nil {
			log.Fatal(err)
		}

		var keyStr string
		for k, t := range tasks {
			if k+1 == int(arg) {
				keyStr = t.Key
			}
		}


		key := []byte(keyStr)
		value, err := s.GetTask(key)
		if err != nil {
			log.Fatal(err)
		}

		if value == nil {
			fmt.Println("Task has not been found.")
			return
		}

		if err = s.DeleteTask(key); err != nil {
			log.Fatal(err)
		}

		if err = s.MarkDone(key, value); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Task has been marked done.")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
