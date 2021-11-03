package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var alphabet = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
}

func namePoints(slice []string, value string)int {
	var sum int
	for _, val := range value {
		letter := string(val)

		for k, v := range slice {
			if v == letter {
				sum += k+1
			}
		}
	}
	return sum
}

func TrimExtra(str string) string {
	res := strings.Trim(str, " ")
	res = strings.Trim(res, "\t")
	res = strings.Trim(res, "\"")
	return res
}

func readStr()[]string{
	file, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var sliceStr []string
	for scanner.Scan() {
		lineStr := scanner.Text()
		for _, v := range strings.Split(lineStr, ","){
			sliceStr = append(sliceStr,TrimExtra(v))
		}
	}
	return sliceStr
}

func findNumberInSortedSlice(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i+1
		}
	}
	return 0
}

func main() {
	var name string
	nameUpper := strings.ToUpper(name)
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)

	sliceOfNames := readStr()
	sort.Strings(sliceOfNames)
	num1 := findNumberInSortedSlice(sliceOfNames,nameUpper)
	num2 := namePoints(alphabet,nameUpper)
	result := num1 * num2
	fmt.Sprintf("%s получает %d очков",nameUpper, result )
}