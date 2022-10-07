package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput string

	fmt.Print("Enter any string: ")
	fmt.Scan(&userInput)

	newString := strings.ToUpper(userInput)
	a := "A"

	if newString[0:1] == "I" && newString[len(newString)-1] == 'N' && strings.Contains(newString, a) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
