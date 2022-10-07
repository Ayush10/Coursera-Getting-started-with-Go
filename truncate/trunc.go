package main

import "fmt"

func main() {
	var floatingNumber float64

	fmt.Print("Type a flaoting point number: ")
	fmt.Scan(&floatingNumber)
	var integerNumber int = int(floatingNumber)
	fmt.Println("Your truncated number is: ", integerNumber)
}
