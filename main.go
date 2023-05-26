package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func main()  {

	reader =bufio.NewReader(os.Stdin)
	fmt.Println("Hello")
	fmt.Print("-> ")

	userInput, _ := reader.ReadString('\n')
	userInput =strings.Replace(userInput,"\r\n","",-1)
	userInput =strings.Replace(userInput,"\n","",-1)

	fmt.Println("Your name is",userInput)
	
}