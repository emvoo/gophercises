package main

import (
		"unicode"
	"strings"
	"bufio"
	"os"
	"fmt"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your string")
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected occurred: %v, try again.", err)
		return
	}

	if len(inputStr) == 0 {
		fmt.Println("Your input must contain at least one character.")
		return
	}

	fmt.Println("Please enter your rotation value")
	var rotation int
	_, err = fmt.Scanf("%d", &rotation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected occurred: %v, try again.", err)
		return
	}

	var ciphered string
	for _, runeValue := range inputStr {
		// if not letter iterate to next character
		if !unicode.IsLetter(runeValue) {
			ciphered += string(runeValue)
			continue
		}

		// check index in the alphabet
		index := strings.Index(alphabet, string(runeValue))
		if (rotation + index) > len(alphabet) - 1 {
			index = index - len(alphabet)
		}
		ciphered += string(alphabet[index+rotation])
	}

	fmt.Println(ciphered)
}
