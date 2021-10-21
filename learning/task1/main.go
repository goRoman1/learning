package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func drawField(width, height int){
	symbol := "*"
	blank := " "
	var sliceOfStars []string
//	joinedSlice := strings.Join(sliceOfStars, blank) // slice becomes a string, but it doesn't work if use it as variable outside the Print
	for i := 0; i < width; i++{
		sliceOfStars = append(sliceOfStars, symbol)
	}
	//	fmt.Println(sliceOfStars)
	for i :=0; i < height; i++{
		if i%2 == 0 {
			fmt.Printf("%s%s\n",strings.Join(sliceOfStars,blank), blank)
		}else{
			fmt.Printf("%s%s\n",blank,strings.Join(sliceOfStars,blank))
		}
	}
}


func main() {
	argsArray := os.Args[1:]
	if len(argsArray) != 2 {
		fmt.Println("Должно быть два аргумента")
	}
	width,err := strconv.Atoi(argsArray[0])
	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(argsArray[1])
	if err != nil {
		log.Fatal(err)
	}

	drawField(width,height)
}
