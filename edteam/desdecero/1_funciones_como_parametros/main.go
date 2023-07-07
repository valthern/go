package main

import "fmt"

func main() {
	/* 	nums := []int{1, 4, 70, 5, 67, 90, 2}

	   	result := filter(nums, func(num int) bool {
	   		if num >= 10 {
	   			return true
	   		}
	   		return false
	   	})

	   	fmt.Println(result) */

	x := hello("Omar")("Martínez")
	fmt.Println(x)
}

func hello(name string) func(string) string {
	return func(text string) string {
		return "Hello " + name + " " + text
	}
}

/* func filter(nums []int, callback func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}

	return result
}
*/
