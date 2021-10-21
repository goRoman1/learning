package main

import (
	"errors"
	"fmt"
	"log"
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

func GetConsoleInput()(string){
	var inputResult string
	_, err := fmt.Scanf("%s\n", &inputResult)
	if err != nil {
		panic(err)
	}
	return inputResult
}

func ParseToFloat(str string) (float64, bool) {
	res, err := strconv.ParseFloat(TrimTabsAndSpaces(str), 64)
	if err != nil {
		return 0, false
	}
	return res, true
}

func CanBeTriangle(sideA, sideB, sideC float64) bool {
	return sideA+sideB > sideC && sideB+sideC > sideA && sideC+sideA > sideB
}

func TrimTabsAndSpaces(str string) string {
	res := strings.Trim(str, " ")
	res = strings.Trim(res, "\t")
	return res
}

func ParseInputStr(inputStr string)(string, float64, float64, float64, error){
	foundParams := strings.Split(inputStr, ",")
	if len(foundParams) != 4 {
		return "", 0, 0, 0, errors.New("number of received parameters doesn't equal to 4")
	}
	name :=  TrimTabsAndSpaces(foundParams[0])
	sideA, okA := ParseToFloat(foundParams[1])
	sideB, okB := ParseToFloat(foundParams[2])
	sideC, okC := ParseToFloat(foundParams[3])

	if !okA || !okB || !okC {
		return "", 0, 0, 0, errors.New("cannot parse one ore more side sizes")
	}

	return name, sideA, sideB, sideC, nil
}

func TriangleInitManually() (string, float64, float64, float64) {
	fmt.Println("Input triangle info: <name>,<sideA>,<sideB>,<sideC>")

	inputResult := GetConsoleInput()
	fmt.Println("input --", inputResult)

	name, sideA, sideB, sideC, err := ParseInputStr(inputResult)
	if err != nil {
		log.Println(err, name,sideA,sideB,sideC)
	}

	if !CanBeTriangle(sideA, sideB, sideC) {
		fmt.Println("The sum of any two parties must be greater than the third")
		panic(err)
	}

	return name, sideA, sideB, sideC
}

func TriangleArea(sideA, sideB, sideC float64)float64{
	p := (sideA + sideB + sideC)/2
	sqrtS := p*(p-sideA)*(p-sideB)*(p-sideC)
	S := math.Sqrt(sqrtS)
	return S
}


func sortTriangles(m map[string]float64) []keyValue {
	newSortedMap := make([]keyValue, 0, len(m))
	for k, v := range m {
		newSortedMap = append(newSortedMap, keyValue{k, v})
	}

	sort.Slice(newSortedMap, func(i, j int) bool {
		return newSortedMap[i].value < newSortedMap[j].value
	})
	return newSortedMap
}

func continueConfirmation()bool{
	var ask string
	fmt.Println("Type 'y' or 'yes' to enter another triangle")
	_, err := fmt.Scan(&ask)
	if err != nil {
		panic(err)
	}
	if strings.ToLower(ask) == "y" || strings.ToLower(ask) == "yes" {
		return true
	} else {
		return false
	}
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
	fmt.Println(sortTriangles(MapUnsorted))
}
