package main

import "fmt"

func main() {
	//sum := 0
	var number int
	result := 1
	fmt.Printf("Enter the number whose factorial you want to calculate\n")
	fmt.Scan(&number)
	for i := number; i >= 1; i-- {
		result *= i
	}
	fmt.Printf("Factorial of %v is %v", number, result)
}
