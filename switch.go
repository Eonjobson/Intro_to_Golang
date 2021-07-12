package main

import "fmt"

func main() {
	colour := "run"
	switch colour {
	case "red":
		fmt.Println("colour is red")
	case "blue":
		fmt.Println("colour is blue")
	default:
		fmt.Println("colour is neither blue nor red")
	}

}
