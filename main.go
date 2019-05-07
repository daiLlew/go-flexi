package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const hhMMFormat = "1504"

var startTime time.Time
var timeWorked time.Time

func main() {
	sc := bufio.NewScanner(os.Stdin)

	initialized := false
	fmt.Println("enter start time - format hh:mm")

	for sc.Scan() {
		in := sc.Text()

		if in == "q" {
			quit()
		}

		if in == "done" {
			finish(timeWorked, startTime)
		}

		start, end := parseInput(in)

		if !initialized {
			startTime = start
			timeWorked = start
			initialized = true
		}

		incrementTotal(start, end)
	}
}

func parseInput(input string) (time.Time, time.Time) {
	times := strings.Split(input, ":")

	fmt.Printf("times: %v\n", times)

	startStr := strings.TrimSpace(times[0])
	endStr := strings.TrimSpace(times[1])

	return parseTime(startStr), parseTime(endStr)
}

func incrementTotal(start time.Time, end time.Time) {
	diffMins := end.Sub(start).Minutes()
	fmt.Printf("diffMins added %f\n", diffMins)
	timeWorked = timeWorked.Add(time.Minute * time.Duration(diffMins))
}

func finish(end time.Time, start time.Time) {
	fmt.Println(end.Sub(start))
	quit()
}

func parseTime(val string) time.Time {
	t, err := time.Parse(hhMMFormat, val)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	return t
}

func quit() {
	fmt.Println("goodbye")
	os.Exit(0)
}
