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
	fmt.Println("Your name is " + userName + " and your age is", age) // + sign don't add extra space, but it is slower as compare to the above. It takes more space also
//and we can't use + with int and string(i.e. two different types)

	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.",userName,age)) // this is must faster
	fmt.Printf("Your name is %s. You are %d years old.\n",userName,age) // better

}

func readSting(s string) string  {

	for { 
		fmt.Println(s)
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput =strings.Replace(userInput,"\r\n","",-1)
		userInput =strings.Replace(userInput,"\n","",-1)

		if userInput == ""{
			fmt.Println("Pls enter yor name")
		} else{
			return userInput
		}
	}
}

func readInt(s string) int  {
	for {
		fmt.Println(s)
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput =strings.Replace(userInput,"\r\n","",-1)
		userInput =strings.Replace(userInput,"\n","",-1)

		num, err:=strconv.Atoi(userInput)// conert string into number

		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}