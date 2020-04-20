package main

import "fmt"
func testEq(a, b []string) bool {

    // If one is nil, the other must \ be nil.
    if (a == nil) != (b == nil) { 
        return false; 
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}
func main() {
  var strSlice = []string{"India", "Canada", "Japan","USA","China"}
  fmt.Printf("Length of slice \tLen: %v\n", len(strSlice))
  fmt.Println(strSlice)
  fmt.Println("Updating index value")
  strSlice[4] = "Germany"
  fmt.Println(strSlice)
  fmt.Println("Copy slice")
  slc3 := make([]string, 5)
  copy_slice := copy(slc3, strSlice) 
  fmt.Println("\nSlice:", slc3) 
  fmt.Println("Total number of elements copied:", copy_slice) 
  slc3[4] = "Russia"
  fmt.Println(slc3)
  fmt.Println(testEq(slc3,strSlice))

}