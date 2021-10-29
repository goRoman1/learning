package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getArgs()(int, int){
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
	return width,height
}

func makeLine(width int)string{
	symbol := "*"
	blank := "_"

	var sliceOfStars []string
	for i := 0; i < width; i++{
		sliceOfStars = append(sliceOfStars, symbol)
	}
	line := strings.Join(sliceOfStars,blank)
	return line
}

func drawField(height int, line string)string{
	blank := "_"
	var board string
	for i :=0; i < height; i++{
		if i%2 == 0 {
			board += line + blank + "\n"
		}else{
			board += blank + line  + "\n"
		}
	}
	return board
}



func main() {
	width,height := getArgs()
	line := makeLine(width)
	fmt.Println(drawField(height,line))
}