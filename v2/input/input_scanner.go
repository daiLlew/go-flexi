package input

import (
	"bufio"
	"fmt"
	"os"
)

type Terminal struct {
	scanner        *bufio.Scanner
	inputReceived  chan string
	inputProcessed chan bool
	exitNow        chan bool
}

func New() *Terminal {
	return &Terminal{
		scanner:        bufio.NewScanner(os.Stdin),
		inputReceived:  make(chan string, 0),
		inputProcessed: make(chan bool, 0),
		exitNow:        make(chan bool, 0),
	}
}

func (t *Terminal) WaitForInput() {
	go func() {
		for t.scanner.Scan() {
			t.inputReceived <- t.scanner.Text()

			<-t.inputProcessed
			fmt.Println("input has been processed continuing")
		}
	}()
}

func (t *Terminal) Received() chan string {
	return t.inputReceived
}

func (t *Terminal) Next() {
	t.inputProcessed <- true
}

func (t *Terminal) Close() {
	close(t.inputProcessed)
	close(t.inputReceived)
	close(t.exitNow)
}
