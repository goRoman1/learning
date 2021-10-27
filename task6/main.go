package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
)

/*
func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}*/

func toIntSlice(str string)[]int{
	slice := strings.Split(str,"")
	var intSlice []int
	for i := range slice {
		elem := slice[i]
		number,err := strconv.Atoi(elem)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, number)
	}
	return intSlice
}

func isPiter(slice []int)bool{
	var odd, even int
	if len(slice) == 6{
		for i := range slice{
//			fmt.Println(i," -i", slice[i],"- slice[i]" )
			if i%2 == 0 {
				even += slice[i]
			}else if i%2 != 0{
				odd += slice[i]
			}
		}
	}
//	fmt.Println(odd, even)
	if odd == even {
		return true
	}
	return false
}



func isMoskow(slice []int)bool{
	var start, fin int
	if len(slice) == 6{
		start = slice[0] + slice[1] + slice[2]
		fin = slice[3] + slice[4] + slice[5]
	}else {
		fmt.Println("len is not 6")
	}
//	fmt.Println(start, fin)
	if start == fin {
		return true
	}
	return false
}

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}


func main(){
	var amount int
	var algorithmPath string
	fmt.Print("Enter path to file with algorithm: ")
	_, err :=fmt.Scan(&algorithmPath)
	if err != nil {
		panic(err)
	}

	strSlice :=  parseFile(algorithmPath)
	if strings.ToLower(strSlice[0]) == "moscow" {
		for _, v := range strSlice[1:]{
			if isMoskow(toIntSlice(v)) {
				amount++
			}
		}
		fmt.Println(amount, " numbers is Moscow")
	}else if strings.ToLower(strSlice[0]) == "piter"  {
		for _, v := range strSlice[1:]{
			if isPiter(toIntSlice(v)){
				amount++
			}
		}
		fmt.Println(amount, " numbers is Piter")
	}

}
