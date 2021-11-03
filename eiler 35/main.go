package main

import (
	"fmt"
	"math"
)

var primes []bool

func findPrimesUpto(n int) []int{
	var result []int
	primes = make([]bool, n+1)
	for i := range primes{
		primes[i] = true
	}
	primes[0] = false
	primes[1] = false
	for i:=2; i*i<=n; i++{
		if primes[i] {
			for j:=i*i; j<=n; j+=i{
				primes[j] = false
			}
		}
	}
	for k, v := range primes{
		if v {
			result = append(result, k)
		}
	}
	return result
}

func leaveCyclic(primeArray []int)[]int{
	result := make([]int, 0, len(primeArray))
	for _, val := range primeArray{
		digits:= digitsFromNumber(val)
		correct := true
		//var checkedNumbers = []int{val}
		for i:=1; i<len(digits); i++{
			newdigits := append(digits[i:], digits[:i]...)
			newnumber := numberFromDigits(newdigits)
			if !primes[newnumber] {
				correct = false
			}
			//checkedNumbers = append(checkedNumbers, newnumber)
		}
		if correct {
			//result = append(result, checkedNumbers...)
			result = append(result, val)
		}
	}

	return result
}

func digitsFromNumber(num int) []int {
	var result []int

	for num>0 {
		digit := num % 10
		num = num / 10
		result = append([]int{digit}, result...)
	}

	return result
}

func numberFromDigits(digits []int) int {
	var result int

	for i, v := range digits{
		result += v * int(math.Pow10(len(digits)-i-1))
	}

	return result
}

func main()  {
	N := 1000000
//	fmt.Println(findPrimesUpto(N))

	result := leaveCyclic(findPrimesUpto(N))
	fmt.Println(result)
	fmt.Println(len(result))


}