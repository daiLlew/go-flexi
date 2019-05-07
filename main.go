package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	isNew := true
	fmt.Println("enter start time - format hh:mm")

	var err error
	var start time.Time
	var end time.Time
	for sc.Scan() {
		in := sc.Text()

		if in == "q" {
			fmt.Println("goodbye")
			os.Exit(0)
		}

		if in == "done" {
			fmt.Println(end.Sub(start))
			os.Exit(0)
		}

		if isNew {
			start, err = time.Parse("15:04", in)
			if err != nil {
				panic(err)
				os.Exit(1)
			}
			end = start

			fmt.Printf("started: %v\n", start)
			fmt.Printf("ended: %v\n", end)
			isNew = false
		} else {
			timeAdded, err := time.Parse("15:04", in)
			if err != nil {
				panic(err)
				os.Exit(1)
			}

			mins := timeAdded.Sub(start).Minutes()
			fmt.Printf("mins added %f", mins)
			end = end.Add(time.Minute * time.Duration(mins))
			fmt.Println(end)
		}

	}
}
