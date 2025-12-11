//Author: Ethan White
//Version: 1.0.0
//Date: 2025-12-11
//Fileoverview: This program will compleate the mathmatical operation requested by the user

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select an operation:")
	fmt.Println("a. add")
	fmt.Println("b. subtract")
	fmt.Println("c. multiply")
	fmt.Println("d. divide")
	fmt.Println("e. absolute value")
	fmt.Println("f. round")
	fmt.Println("g. raise to an exponent")
	fmt.Println("h. square root")

	fmt.Print("Enter choice: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(strings.ToLower(choice))

	var num1, num2 float64

	if choice == "a" { // add
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		fmt.Println(num1, "+", num2, "=", num1+num2)
	} else if choice == "b" { // subtract
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		fmt.Println(num1, "-", num2, "=", num1-num2)
	} else if choice == "c" { // multiply
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		fmt.Println(num1, "*", num2, "=", num1*num2)
	} else if choice == "d" { // divide
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		fmt.Println(num1, "/", num2, "=", num1/num2)
	} else if choice == "e" { // absolute value
		fmt.Print("Enter number: ")
		fmt.Scanln(&num1)
		fmt.Println("|", num1, "| =", math.Abs(num1))
	} else if choice == "f" { // round
		fmt.Print("Enter number: ")
		fmt.Scanln(&num1)
		fmt.Println("round(", num1, ") =", math.Round(num1))
	} else if choice == "g" { // exponent
		fmt.Print("Enter base: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter exponent: ")
		fmt.Scanln(&num2)
		fmt.Println(num1, "^", num2, "=", math.Pow(num1, num2))
	} else if choice == "h" { // square root
		fmt.Print("Enter number: ")
		fmt.Scanln(&num1)
		fmt.Println("sqrt(", num1, ") =", math.Sqrt(num1))
	} else {
		fmt.Println("Invalid choice.")
	}
	fmt.Println("\nDone.")
}
