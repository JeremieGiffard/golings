// anonymous functions3
// Make me compile!

package main

import "fmt"

var index = 1

func updateStatus() func() string {

	orderStatus := map[int]string{
		1: "TO DO",
		2: "DOING",
		3: "DONE",
	}

	return func() string {
		index++
		return orderStatus[index]
	}
}

func main() {
	anonymous_func := updateStatus()
	var status string

	status = anonymous_func()
	status = anonymous_func()

	if status == "DONE" {
		fmt.Println("Good Job!")
	} else {
		fmt.Println(status)
		panic("To complete the challenge the status must be DONE")
	}
}
