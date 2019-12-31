package main

import "fmt"

func main() {
	a := make(chan string)
	go func() {
		for {
			if item, ok := <-a; ok {
				fmt.Println(item)
			}
		}
	}()
	a <- "data1"
	fmt.Println("send data1")
	a <- "data2"
	fmt.Println("send data2")
	a <- "data3"
	fmt.Println("send data3")
	close(a)
}