package main

import "fmt"

func main() {
	arr := make([]int, 0)
	arr = append(arr, 1)

	fmt.Println(arr)
	test(arr)
	fmt.Println(arr)

}

func test(arr []int) {
	arr = append(arr, 1)
}
