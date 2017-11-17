package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool, 1)

	select {
	case msg := <- messages:
		fmt.Println("Received message", msg)
	default:
		fmt.Println("No message received") 
	}

	msg := "Hello World"
	select {
	case messages <- msg:
		fmt.Println("Submitted message", msg)
	default:
		fmt.Println("No message submitted")
	}

	select {
	case msg := <-messages:
			fmt.Println("Received message", msg)
	case sig := <-signals:
			fmt.Println("Received signal", sig)
	default:
			fmt.Println("No activity")
	}
}