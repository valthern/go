package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 6, 3, 6))
}

func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}
	return total
}
