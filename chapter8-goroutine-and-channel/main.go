package main

import "fmt"

func main() {
	channel := make(chan int, 5)
	for i := 0; i < 6; i++ {
		channel <- i
	}
	for len(channel) > 0 {
		fmt.Printf("value of :%d,length of chan:%d", <-channel, len(channel))
		fmt.Println()
	}
	close(channel)
}
