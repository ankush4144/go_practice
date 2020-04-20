// Go program to illustrate how to 
// create and initialize maps 
package main 

import "fmt"

func main() { 

	map_1 := map[int]string{ 
	
			1: "Ram", 
			2: "Sita", 
			3: "Laxman", 
			4: "Jay", 
			5: "Geeta", 
	} 
	
	fmt.Println("Map: ", map_1) 
  map_1[2]="Hanuman"
  fmt.Println("Map: ", map_1)
  delete(map_1, 4)
  fmt.Println("map:", map_1)
} 