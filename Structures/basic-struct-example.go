package main

import (
	"fmt"
)

type Umuntu struct {
	name    string
	surname string
	age     int
	isMale  bool
}

func main() {
	person1 := Umuntu{name: "Dineo", surname: "Mokoena", age: 26, isMale: false}
	person2 := Umuntu{"Bheki", "Sewawa", 35, true}

	fmt.Println("Person 1 name :",person1.name)
	fmt.Println("Person 1 information :", person1)


	fmt.Println("Person 2 name :",person2.name)
	fmt.Println("Person 2 information :", person2)
}
