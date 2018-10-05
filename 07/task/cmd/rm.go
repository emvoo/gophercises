package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"gophercises/07/task/store"
	"fmt"
)

var rmCmd = &cobra.Command{
	Use:"rm",
	Short:"Removes the task from your TODO list.",
	Args:cobra.MaximumNArgs(1),
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

		if err = s.DeleteTask(key); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Task has been successfully deleted.")
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}