package main

import "fmt"

func main(){
	fmt.Println(Max([]int{2,1,6,3,8,32,16}))
}

func Max(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func MaxIndex(numbers []int) int {
	var max int
	var index int

	for i, number := range numbers {
		if number > max {
			max = number
			index = i
		}
	}
	return index
}