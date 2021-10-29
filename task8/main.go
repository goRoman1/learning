package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
/*
func fibRecursion(n int) int {
	if n < 2 {
		return 1
	}
	return fibRecursion(n-2) + fibRecursion(n-1)
}
 */

//количество чисел в промежутке
func countFibs(low, high int) []int {
	var num1, num2, num3 = 0, 1, 1
	var numSlice []int
	//0, 1, 1, 2, З, 5, 8, 1З, 21, З4, 55, 89, 144
//	var result int
	for num1 <= high{
		if num1 > low{
//			result ++  //количество чичел в промежутке
			numSlice = append(numSlice, num1)
//			fmt.Println(num1)
		}
		num1 = num2
		num2 = num3
		num3 = num1 + num2
	}
	return  numSlice
}


func makeString(values []int) string {
	var valuesText []string
	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	result := strings.Join(valuesText, ",")
	return result
}

func main() {
	argsArray := os.Args[1:]
	if len(argsArray) != 2 {
		fmt.Println("Должно быть два аргумента")
	}
	start,err := strconv.Atoi(argsArray[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(argsArray[1])
	if err != nil {
		log.Fatal(err)
	}

	var sliceOfInts = countFibs(start,end)
	var StringFib = makeString(sliceOfInts)
	fmt.Println(StringFib)
}