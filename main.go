package main

import "fmt"

func main () {

	for i:=0; i<3 ;i++ {
		fmt.Print("i: ",i)
		for j:=0; j<3;j++ {
			fmt.Print("    j: ",j)
		}
		fmt.Println()
 	}
}