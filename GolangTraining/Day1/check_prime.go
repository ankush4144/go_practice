package main

import "fmt"

func checkprime(number int) bool {
	if number == 1 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	//sum := 0
	var number int
	fmt.Printf("Enter the number to check for prime\n")
	fmt.Scan(&number)
	if checkprime(number) {
		fmt.Printf("%v is prime", number)
	} else {
		fmt.Printf("%v is not prime", number)
	}
}
