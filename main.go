package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/ONSdigital/log.go/log"
)

const hhMMFormat = "1504"

var (
	startTime   time.Time
	timeWorked  time.Time
	initialized = false
	w *tabwriter.Writer
)

func Init() {
	log.Namespace = "go-flexi"
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	w = tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent|tabwriter.Debug)
	fmt.Fprintln(w, "Enter times\tFormat <hh:mm hh:mm>, <hh:mm hh:mm>, ... <hh:mm hh:mm>")
	fmt.Fprintln(w, fmt.Sprintf("Example \t %q", "0800 0900, 1025 1100, 1200 1650"))
	fmt.Println()
	w.Flush()

	for sc.Scan() {
		input := sc.Text()

		if input == "q" {
			quit()
		}

		if input == "done" {
			fmt.Println()
			fmt.Fprintln(w, fmt.Sprintf("Total time \t %v", timeWorked.Sub(startTime)))
			w.Flush()
			fmt.Println()
			quit()
		}

		process(input)
	}
}

func process(input string) {
	periods := strings.Split(input, ",")

	var totalMins float64 = 0
	fmt.Fprintln(w, "Period\t Duration (minutes)")

	for _, p := range periods {
		p := strings.TrimSpace(p)

		times := strings.Split(p, " ")
		startStr := strings.TrimSpace(times[0])
		endStr := strings.TrimSpace(times[1])

		start := parseTime(startStr)
		end := parseTime(endStr)

		if !initialized {
			startTime = start
			timeWorked = end
			initialized = true
		}

		mins := end.Sub(start).Minutes()
		fmt.Fprintln(w, fmt.Sprintf("%s - %s\t %f", startStr, endStr, mins))

		totalMins += mins
	}

	timeWorked = startTime.Add(time.Minute * time.Duration(totalMins))

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
	os.Exit(0)
}
