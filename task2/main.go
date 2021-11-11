package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Есть два конверта со сторонами (a,b) и (c,d) определить, можно ли один конверт вложить в другой. Программа должна обрабатывать
ввод чисел с плавающей точкой. Программа спрашивает у пользователя размеры конвертов по одному параметру за раз. После каждого
подсчёта программа спрашивает у пользователя хочет ли он продолжить. Если пользователь ответит “y” или “yes” (без учёта регистра),
программа продолжает работу сначала, в противном случае – завершает выполнение.
 */

type Envelope struct {
	height, width float64
}

type EnvelopeI interface {
	maxSide() float64
	minSide() float64
	CanBeNested(envelopeToFit Envelope) bool
}

func envelopeInitialisation()Envelope{
	var sides []float64
	for i:=0; i < 2; i++{
		fmt.Println("Enter side ", i+1)
		var sideInput string
		_, err := fmt.Scanln(&sideInput)
		if err != nil {
			fmt.Println("An error occurred on iteration:", i+1 , "You did not enter a value")
			i--
			continue
		}

		floatParam, err := strconv.ParseFloat(sideInput, 64)
		if err != nil || floatParam <= 0{
			fmt.Println("Value should contain only numbers greater than 0")
			i--
			continue
		}
		sides = append(sides,floatParam)
	}
	return Envelope{sides[0], sides[1]}
}

func (env Envelope) maxSide() float64 {
	if env.height > env.width {
		return env.height
	}
	return env.width
}

func (env Envelope) minSide() float64 {
	if env.height < env.width {
		return env.height
	}
	return env.width
}

func (env Envelope) CanBeNested(envelopeToFit Envelope) bool {
	if env.maxSide() < envelopeToFit.maxSide() && env.minSide() < envelopeToFit.minSide() {
		return true
	}
	return false
}

func checkNesting(envelope1, envelope2 Envelope) string {
	if envelope1.CanBeNested(envelope2){
		return "First can be nested into Second"
	}else if envelope2.CanBeNested(envelope1){
		return "Second can be nested into First"
	}
	return "Envelopes cannot be nested to each other"
}

func askForContinuing()string{
	fmt.Println("Type 'y' or 'yes' to continue")
	var ask string
	fmt.Scanln(&ask)
	return ask
}

func checkIfContinuing(answer string)bool{
	answer = strings.TrimSpace(strings.ToLower(answer))
	if answer == "yes" || answer == "y" {
		return true
	}
	fmt.Println("Program exited")
	return false
}

func main() {
	for{
		fmt.Println("Enter first envelope parameters")
		envelope1 := envelopeInitialisation()

		fmt.Println("Enter second envelope parameters")
		envelope2 := envelopeInitialisation()

		fmt.Println(checkNesting(envelope1, envelope2))
		if !checkIfContinuing(askForContinuing()){
			break
		}
	}
}