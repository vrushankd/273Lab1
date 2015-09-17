package main

import "fmt"

func fibonacci(number int) int {
	if (number == 0) || (number == 1) {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
}

func main() {
	var number int
	fmt.Print("Enter the number to find a fibonnaci: ")
	fmt.Scanf("%d", &number)
	if number >= 0 {
		result := fibonacci(number)
		fmt.Printf("The fibonacci of %d is %d", number, result)
	} else {
		fmt.Println("Warning !! Enter a positive number")
	}
}
