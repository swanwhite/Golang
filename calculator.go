package main

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	a      int
	b      int
	method string
)

func calculate(a, b int) int {
	fmt.Println("Which method of calculation? (+ or - or / or *): ")
	fmt.Scanln(&method)

	fmt.Println("Enter the first number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&b)

	if method == "+" {
		addSum := a + b
		return addSum
	} else if method == "-" {
		minusSum := a - b
		return minusSum
	} else if method == "/" {
		divideSum := a / b
		return divideSum
	} else if method == "*" {
		multiplySum := a * b
		return multiplySum
	} else if method != `"+" "-" "/" "*"` {
		color.Red("Error when calculating. Please enter a valid method of calculation.")
	}
	return a

}

func main() {
	for {
		fmt.Println("\nCalculator by chris. 2021")
		fmt.Println("\n--------------------")
		fmt.Println("The result is: ", calculate(a, b))
	}

}
