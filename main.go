package main

import "fmt"



func main() {

	// Sender data.
	err := readEvents(kafkaEventHandler)
	if err != nil {
		fmt.Println(err)
	}
}
