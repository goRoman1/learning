package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var MapUnsorted = make(map[string]float64)

type keyValue struct {
	key   string
	value float64
}

func GetConsoleInput()string{
	var inputResult string
	_, err := fmt.Scanf("%s\n", &inputResult)
	if err != nil {
		panic(err)
	}
//	fmt.Println("input res ----",inputResult)
	return inputResult
}

func ParseToFloat(str string) (float64, bool) {
	res, err := strconv.ParseFloat(TrimTabsAndSpaces(str), 64)
	if err != nil {
		return 0, false
	}
	return res, true
}

func IsTriangle(sideA, sideB, sideC float64) bool {
	return sideA+sideB > sideC && sideB+sideC > sideA && sideC+sideA > sideB
}

func TrimTabsAndSpaces(str string) string {
	res := strings.Trim(str, " ")
	res = strings.Trim(res, "\t")
	res = strings.Trim(res, " ")
	return res
}

func ParseInputStr(inputStr string)(string, float64, float64, float64){
	foundParams := strings.Split(inputStr, ",")
	if len(foundParams) != 4 {
		fmt.Println("number of received parameters doesn't equal to 4")
	}
	name :=  TrimTabsAndSpaces(foundParams[0])
	sideA, okA := ParseToFloat(foundParams[1])
	sideB, okB := ParseToFloat(foundParams[2])
	sideC, okC := ParseToFloat(foundParams[3])

	if !okA || !okB || !okC {
		fmt.Println("cannot parse one ore more side sizes")
	}

	return name, sideA, sideB, sideC
}

func TriangleInitManually() (string, float64, float64, float64) {
	fmt.Println("Input triangle info: <name>,<sideA>,<sideB>,<sideC>")

	inputResult := GetConsoleInput()
	name, sideA, sideB, sideC:= ParseInputStr(inputResult)

	if !IsTriangle(sideA, sideB, sideC) {
		fmt.Println("The sum of any two parties must be greater than the third")
	}

	return name, sideA, sideB, sideC
}

func TriangleArea(sideA, sideB, sideC float64)float64{
	p := (sideA + sideB + sideC)/2
	sqrtS := p*(p-sideA)*(p-sideB)*(p-sideC)
	S := math.Sqrt(sqrtS)
	return S
}

func continueConfirmation()bool{
	var ask string
	fmt.Println("Type 'y' or 'yes' to enter another triangle")
	_, err := fmt.Scanln(&ask)
	if err != nil {
		fmt.Println(err)
	}
//	ask = strings.Trim(ask, "\n")
	if strings.ToLower(ask) == "y" || strings.ToLower(ask) == "yes" {
		return true
	} else {
		return false
	}
}

func triangleSort(m map[string]float64) []keyValue {
	newSortedMap := make([]keyValue, 0, len(m))
	for k, v := range m {
		newSortedMap = append(newSortedMap, keyValue{k, v})
	}

	sort.Slice(newSortedMap, func(i, j int) bool {
		return newSortedMap[i].value < newSortedMap[j].value

	})

	 return newSortedMap
}

func main() {
	for {
		name, sideA, sideB, sideC := TriangleInitManually()
		MapUnsorted[name] = TriangleArea(sideA, sideB, sideC)
		if !continueConfirmation() {
			break
		}
	}
	fmt.Println("=========Triangles list:=============")
	fmt.Println(triangleSort(MapUnsorted))

}
