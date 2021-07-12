package main

import "fmt"

func main() {
	// ids := []int{6, 5, 4, 43, 21}
	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }
	// //not using index
	// for _, id := range ids {
	// 	fmt.Printf("ID: %d\n", id)
	// }
	// in maps
	numbers := map[int]string{1: "one", 2: "two"}
	for k, v := range numbers {
		fmt.Printf("%d:%s\n", k, v)
	}
}
