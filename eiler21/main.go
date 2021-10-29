package main

import "fmt"

func friendlyForNumber(n int) int {
	var a int
	for i := 1; i < n; i ++ {
		if n%i == 0{
			a += i
		}
	}
	return a
}

func sliceOfFriendlyNums(upper int)[]int{
	var friendlySlice []int
	for i := 0; i < upper; i++ {
		sum := friendlyForNumber(i)
		if sum == i {
			continue
		}
		if friendlyForNumber(sum) == i {
			friendlySlice = append(friendlySlice, i)
		}
	}
	return friendlySlice
}

func main (){
	var sum int
	friendlySlice := sliceOfFriendlyNums(10000)
	for i := 0; i < len(friendlySlice); i++ {
		sum += friendlySlice[i]
	}
	fmt.Println(sum)

}

