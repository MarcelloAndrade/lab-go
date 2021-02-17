package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Init..")
	channel := make(chan int)
	fmt.Println("Call Channel..")
	go teste(channel)
	fmt.Println("Back Channel..")

	fmt.Printf("\n")

	fmt.Println("-> Print Values 1..")
	fmt.Println("-> Print Values 1..: ", <-channel, <-channel)

	fmt.Printf("\n")

	fmt.Println("-> Print Values 2..")
	fmt.Println("-> Print Values 2..:", <-channel)
}

func teste(channel chan int) {
	fmt.Println("	Channel..")

	time.Sleep(time.Second)
	fmt.Println("	Set Channel 1 ..")
	channel <- 1
	fmt.Println("	Back Set Channel 1 ..")

	time.Sleep(time.Second)
	fmt.Println("	Set Channel 2 ..")
	channel <- 2
	fmt.Println("	Back Set Channel 2 ..")

	time.Sleep(3 * time.Second)
	fmt.Println("	Set Channel 3 ..")
	channel <- 3
	fmt.Println("	Back Set Channel 3 ..")
}
