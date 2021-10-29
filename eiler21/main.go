package main

import "fmt"

func friendlyForNumber(n int) int {
	var sum int
	for i := 1; i < n; i ++ {
		if n%i == 0{
			sum += i
		}
	}
	return sum
}

func sumInRange(upper int)int{
	var sum int
	for i := 0; i < upper; i++ {
		num := friendlyForNumber(i)
		if num == i {
			continue
		}
		if friendlyForNumber(num) == i {
			sum += i
		}
	}
	return sum
}

func main (){

	fmt.Println(sumInRange(10000))

}

