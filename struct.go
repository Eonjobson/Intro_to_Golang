package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
	gender    string
}

func main() {
	person1 := Person{"Rohin", "Joshi", 19, "M"}
	fmt.Println(person1)
}
