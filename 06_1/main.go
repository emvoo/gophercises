package main

import (
	"unicode"
	"fmt"
	"bufio"
	"os"
	"log"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	count := 1
	for _, runeValue := range s {
		if unicode.IsUpper(runeValue) {
			count++
		}
	}
	fmt.Println(count)
}
