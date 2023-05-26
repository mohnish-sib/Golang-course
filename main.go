package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main()  {

	reader =bufio.NewReader(os.Stdin)

	userName := readSting("Hello")
	age:=readInt("Your age?")

	fmt.Println("Your name is",userName, "and your age is", age)
	
}

func readSting(s string) string  {
	fmt.Println(s)
	fmt.Print("-> ")

	userInput, _ := reader.ReadString('\n')
	userInput =strings.Replace(userInput,"\r\n","",-1)
	userInput =strings.Replace(userInput,"\n","",-1)

	return userInput
}

func readInt(s string) int  {
	fmt.Println(s)
	fmt.Print("-> ")

	userInput, _ := reader.ReadString('\n')
	userInput =strings.Replace(userInput,"\r\n","",-1)
	userInput =strings.Replace(userInput,"\n","",-1)

	num, err:=strconv.Atoi(userInput)// conert string into number

	if err != nil{
		fmt.Println("Please enter a whole number")
	}

	return num
}