package main

import (
	"fmt"
)

func main() {

	var x int
	fmt.Print("enter a number between 1-3: ")

	fmt.Scan(&x)
	if x == 1 {
		fmt.Println("Hello World!")
	} else if x == 2 {
		fmt.Println("Hello Dante!")
	} else if x == 3 {
		fmt.Println("Hello Dale!")
	}
}
