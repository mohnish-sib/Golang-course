package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName 		string
	Age 			int
	FavouriteNumber float64
	OwnsADog 		bool
}

func main()  {

	reader =bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readSting("Hello")
	user.Age=readInt("Your age?")
	user.FavouriteNumber =readFloat("Enter your fav no.-")
	user.OwnsADog = readBool("Do you own a dog (y/n)?")

	// fmt.Println("Your name is",userName, "and your age is", age)
	// fmt.Println("Your name is " + userName + " and your age is", age) // + sign don't add extra space, but it is slower as compare to the above. It takes more space also
	//and we can't use + with int and string(i.e. two different types)

	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.",userName,age)) // this is must faster
	fmt.Printf("Your name is %s. You are %d years old. And your fav number is %.2f. Owns a dog: %t. \n",
		user.UserName, 
		user.Age, 
		user.FavouriteNumber, 
		user.OwnsADog) // better

}

func readSting(s string) string  {

	for { 
		fmt.Println(s)
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput =strings.Replace(userInput,"\r\n","",-1)
		userInput =strings.Replace(userInput,"\n","",-1)

		if userInput == "" {
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

func readFloat(s string) float64  {
	for {
		fmt.Println(s)
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput =strings.Replace(userInput,"\r\n","",-1)
		userInput =strings.Replace(userInput,"\n","",-1)

		num, err:=strconv.ParseFloat(userInput,64)// conert string into float

		if err != nil {
			fmt.Println("Please enter a correct number")
		} else {
			return num
		}
	}
}

func readBool(s string) bool  {
	for {
		fmt.Println(s)
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput =strings.Replace(userInput,"\r\n","",-1)
		userInput =strings.Replace(userInput,"\n","",-1)

		if userInput != "y" && userInput !="n" {
			fmt.Println("Please enter y or n")
		} else {
			return userInput == "y"
		}
	}
}