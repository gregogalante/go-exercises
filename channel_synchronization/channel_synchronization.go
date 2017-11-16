package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working on gorountime...")
	time.Sleep(time.Second)
	fmt.Println("Done on gorountime!")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	
	go worker(done)

	fmt.Println("Do new things on main...")

	<-done
}