package main

import "fmt"

func main() {
	number := 1
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", number)
			number += 1
		}
		fmt.Println()
	}
}
