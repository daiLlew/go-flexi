package main

import (
	"fmt"
	"os"

	"github.com/daiLlew/go-flexi/v2/input"
)

type TimeCalculator struct {
	input *input.Terminal
}

func main() {
	timeCalculator := &TimeCalculator{
		input: input.New(),
	}

	timeCalculator.start()
}

func (calc *TimeCalculator) start() {
	calc.input.WaitForInput()

	done := false
	for !done {
		select {
		case input := <-calc.input.Received():
			calc.inputReceived(input)
		}
	}
}

func (calc *TimeCalculator) inputReceived(input string) {
	if input == "q" {
		calc.exit()
	}

	fmt.Printf("input received: %q\n", input)
	calc.input.Next()
}

func (calc *TimeCalculator) exit() {
	fmt.Println("quitting app - goodbye")
	calc.input.Close()
	os.Exit(0)
}
