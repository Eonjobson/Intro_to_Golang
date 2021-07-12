package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("enter first number")
	fmt.Scanln(&num1, &num2)
	add, sub, div, mult := calc(num1, num2)
	fmt.Println("Sum is ", add)
	fmt.Println("difference is ", sub)
	fmt.Println("quotient is ", div)
	fmt.Println("product is ", mult)

}

func calc(x, y int) (add, sub, div, mult float32) {
	add = float32(x + y)
	sub = float32(x - y)
	div = float32(x / y)
	mult = float32(x * y)
	return

}
