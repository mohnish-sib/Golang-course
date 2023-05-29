package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"myapp/mylogger"
)

func main() {
	// read input from the user 5 times and write it to a log

	// first, create the reader, and a channel of strings
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	// start the ListenForLog function in the mylogger package
	// as a goroutine, so it runs in the background
	go mylogger.ListenForLog(ch)

	// give some instructions
	fmt.Println("Enter something")

	// execute a loop five times, using the three component for loop
	for i := 0; i < 5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input

		// wait one second, so the log has time to print to the terminal window
		time.Sleep(time.Second)
	}
}
