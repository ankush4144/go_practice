package main

import (
	"fmt"
	"os"
)

// Student structure
type Student struct {
	firstName, lastName string
	contact_number      string
}

func update_name(p *Student) {
	p.firstName = "TEST"
}

func createFile() {
	// check if file exists
	var _, err = os.Stat("data.txt")

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create("data.txt")
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", "data.txt")
}

func writeFile(data string) {
	var file, err = os.OpenFile("data.txt", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(data)
	if isError(err) {
		return
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func main() {
	stud2 := Student{"Johny", "Garg", "9813456895"}
	fmt.Println("Student 2 info", stud2)

	pts := &stud2

	update_name(&stud2)
	name := fmt.Sprintf("Name : %v", pts.firstName)
	contact := fmt.Sprintf("Contact : %v", pts.contact_number)
	createFile()
	writeFile(name)
	writeFile(contact)
}
