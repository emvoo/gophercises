package main

import (
	"flag"
	"time"
	"os"
	"log"
	"encoding/csv"
	"io"
	"fmt"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "Use this and that.")
	duration := flag.Int("d", 30, "Used to set test duration")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		log.Printf("failed to open file: %v", err)
	}

	r := csv.NewReader(file)

	strngs := [][]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("3: %v", err)
			break
		}

		strngs = append(strngs, record)
	}
	fmt.Println("Please press ENTER to begin and start the counter!")
	fmt.Scanf("\n")
	tDuration := time.Second * time.Duration(*duration)
	timer := time.NewTimer(tDuration)
	var count int
	for _, item := range strngs {
		answerCh := make(chan string)
		go func() {
			fmt.Printf("%s = ", item[0])
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nYou run out of time")
			fmt.Printf("\nYou scored %d out od %d\n", count, len(strngs))
			return
		case answer := <-answerCh:
			if answer == item[1] {
				count++
			}
		}
	}

	fmt.Printf("\nYou scored %d out od %d\n", count, len(strngs))
}
