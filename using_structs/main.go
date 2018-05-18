package main

import "fmt"

type person struct {
	firstName	string
	lastName	string
}

func scanName() string {
	var n string
	fmt.Scanln(&n)
	return n
}

func getNames() person {
	fmt.Print("What is your first name: ")
	firstName := scanName()
	fmt.Print("What is your last name: ")
	lastName := scanName()
	return person{firstName, lastName}
}

func main() {
	fmt.Println(getNames())
}
