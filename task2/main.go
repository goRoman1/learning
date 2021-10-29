package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func setEnvelopParameters() (Envelope, error) {
	height := triangleSideSet("Envelope height:")
	width := triangleSideSet("Envelope width:")
	return Envelope{height, width}, nil
}

func triangleSideSet(str string)float64{
	fmt.Println(str)
	var inp string
	_, err := fmt.Scanln(&inp)
	if err != nil {
		fmt.Println("An error occurred 1:", err)
	}

	floatParam, err := strconv.ParseFloat(inp, 64)
	if err != nil {
		fmt.Println("An error occurred 2:", err)
	}

	if floatParam < 0 {
		fmt.Println("Side length must be longer then 0")
	}
	return floatParam
}

func askForContinuing() bool{
	fmt.Println("Type 'y' or 'yes' to continue")
	var ask string
	_, err := fmt.Scanln(&ask)
	if err != nil {
		return false
	}

	ask = strings.TrimSpace(strings.ToLower(ask))
	if ask == "yes" || ask == "y"{
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

func main() {
	for{
		fmt.Println("Enter first envelope parameters")
		envelope1, err := setEnvelopParameters()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Enter second envelope parameters")
		envelope2, err := setEnvelopParameters()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(checkNesting(envelope1, envelope2))

		if !askForContinuing(){
			break
		}
	}
}