package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func countNums(num float64)[]int{
	var x []int
	sqrtNum := math.Sqrt(num)
	fmt.Println(sqrtNum)
	for i := 0; i < int(sqrtNum); i++{
		x = append(x, i)
	}
	return x
}

func makeSliceOfStrings(values []int){
	var valuesText []string
	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	result := strings.Join(valuesText, ",")
	fmt.Println(result)
}

func main(){
	argsArray := os.Args[1:]
	if len(argsArray) != 1 {
		fmt.Println("Должен быть один аргумент")
	}

	floatParam, _ := strconv.ParseFloat(strings.TrimLeft(argsArray[0], "-"), 64)

	sliceOfInt := countNums(floatParam)
	makeSliceOfStrings(sliceOfInt)

}
