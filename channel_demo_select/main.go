package main

import (
	"fmt"
	"os"
)

var (
	chan1 chan bool
	chan2 chan bool
	quit chan bool
)

func main() {
	chan1 = make(chan bool)
	chan2 = make(chan bool)
	quit = make(chan bool)
	go handler()
	chan1 <- true
	chan2 <- true
	close(chan1)
	<-quit
}

func handler() {
	for {
		select {
		case _, ok := <-chan1:
			if ok {
				fmt.Println("chan1")
			} else {
				fmt.Println("chan1 closed")
				os.Exit(0)
			}
		case <-chan2:
			fmt.Println("chan2")
			quit <- true
		}
	}
}