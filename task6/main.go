package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseFile(path string)[]string{
	var lines []string
	file, err := os.Open(path) //"1.txt"
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err();
		err != nil {
		log.Fatal(err)
	}
	return lines
}

func executeAlgorithm(strSlice []string)string{
	var amount int
	switch strings.ToLower(strSlice[0]) {
	case  "moscow":
		for _, v := range strSlice[1:]{
			if isMoscow(stringToIntSlice(v)) {
				amount++
			}
		}
		return strconv.Itoa(amount)+ " numbers is Moscow"
	case "piter":
		for _, v := range strSlice[1:]{
			if isPiter(stringToIntSlice(v)){
				amount++
			}
		}
		return strconv.Itoa(amount)+ " numbers is Piter"
	case "truepiter":
		for _, v := range strSlice[1:]{
			if isTruePiter(stringToIntSlice(v)){
				amount++
			}
		}
		return strconv.Itoa(amount)+ " numbers is truePiter"
	default:
		return "unknown algorithm"
	}
}

func stringToIntSlice(str string)[]int{
	slice := strings.Split(str,"")
	var intSlice []int
	for _, elem := range slice {
		number,err := strconv.Atoi(elem)
		if err != nil {
			continue
		}
		intSlice = append(intSlice, number)
	}
	return intSlice
}

func isTruePiter(slice []int)bool{
	var odd, even int
	if len(slice) != 6{
		return false
	}

	for i := range slice{
		if i%2 == 0 {
			even += slice[i]
		}else if i%2 != 0{
			odd += slice[i]
		}
	}

	return odd == even
}

func isPiter(slice []int)bool{
	var odd, even int
	if len(slice) != 6{
		return false
	}
	for _ , v:= range slice{
		if v%2 == 0 {
			even += v
		}else if v%2 != 0{
			odd += v
		}
	}
	return odd == even
}

func isMoscow(slice []int)bool{
	var start, fin int
	if len(slice) != 6{
		return false
	}

	start = slice[0] + slice[1] + slice[2]
	fin = slice[3] + slice[4] + slice[5]

	return start == fin
}

func main(){
	var algorithmPath string
	fmt.Print("Enter path to file with algorithm: ")
	_, err :=fmt.Scan(&algorithmPath)
	if err != nil {
		panic(err)
	}

	strSlice :=  parseFile(algorithmPath)
	fmt.Println(executeAlgorithm(strSlice))
}
