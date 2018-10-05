package main

import (
	"gophercises/07/task/cmd"
	"os"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cmd.RootCmd.Execute()
	os.Exit(1)
}
