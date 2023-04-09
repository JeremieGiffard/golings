// structs2
// Make me compile!
package main

import "fmt"

type Contact struct {
	phone string
}
type Person struct {
	// don't just create the phone field here. embed a new struct
	Contact
	name string
	age  int
}

func main() {
	person := Person{name: "John", age: 32, Contact: Contact{phone: "20.15.45"}}
	//person := Person{name: "John", age: 32}
	person.phone = "2145356"
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
