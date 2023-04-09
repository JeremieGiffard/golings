// structs1
// Make me compile!
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var person = new(Person)
	person.name = "truc"
	person.age = 10
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
