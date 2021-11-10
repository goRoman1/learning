package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func envelopeInitialisation()Envelope{
	var sides []float64
	for i:=0; i < 2; i++{
		fmt.Println("Enter side ", i+1)
		var sideInput string
		_, err := fmt.Scanln(&sideInput)
		if err != nil {
			fmt.Println("An error occurred on iteration:", i+1 , err)
			break
		}

		floatParam, err := strconv.ParseFloat(sideInput, 64)
		if err != nil || floatParam <= 0.0{
			fmt.Println("Value should contain only numbers greater than 0")
			i--
			continue
		}
		sides = append(sides,floatParam)
	}
	return Envelope{sides[0], sides[1]}
}

func checkNesting(envelope1, envelope2 Envelope) string {
	if envelope1.CanBeNested(envelope2){
		return "First can be nested into Second"
	}else if envelope2.CanBeNested(envelope1){
		return "Second can be nested into First"
	}
	return "Envelopes cannot be nested to each other"
}

func askUserContinue() bool {
	fmt.Println("\nEnter 'y' or 'yes' if you want to add one more triangle")
	var answer string
	fmt.Scan(&answer)
	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		return true
	}
	return false
}

func userInput(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	fmt.Print(`Enter parameters of triangle in format: <name>, <side>, <side>, <side> need comma between as a delimiter: `)
	var input string
	scanner.Scan()
	input = scanner.Text()
	if input == "" {
		fmt.Println("Must be a value")
		return ""
	}
	return input
}

func main() {
	input := userInput(os.Stdin)

	prepareInput(i)

	fmt.Println("Enter first envelope parameters")
	envelope1 := envelopeInitialisation()
	fmt.Println("Enter second envelope parameters")
	envelope2 := envelopeInitialisation()

	if askUserContinue() {
		main()
		return
	}
	fmt.Println(checkNesting(envelope1, envelope2))


}