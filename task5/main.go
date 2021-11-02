package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var russianUnits = []string{
	"",
	"один",
	"два",
	"три",
	"четыре",
	"пять",
	"шесть",
	"семь",
	"восемь",
	"девять",
}
var russianTeens = []string{
	"десять",
	"одиннадцать",
	"двенадцать",
	"тринадцать",
	"четырнадцать",
	"пятнадцать",
	"шестнадцать",
	"семнадцать",
	"восемнадцать",
	"девятнадцать",
}
var russianTens = []string{
	"",
	"десять",
	"двадцать",
	"тридцать",
	"сорок",
	"пятьдесят",
	"шестьдесят",
	"семьдесят",
	"восемьдесят",
	"девяносто",
}
var russianHundreds = []string{
	"",
	"сто",
	"двести",
	"триста",
	"четыреста",
	"пятьсот",
	"шестьсот",
	"семьсот",
	"восемьсот",
	"девятьсот",
}
var russianMegas = [][]string{
	{"", "", ""},
	{"тысяча", "тысячи", "тысяч"},                    // 10^3
	{"миллион", "миллиона", "миллионов"},             // 10^6
	{"миллиард", "миллиарда", "миллиардов"},          // 10^9
	{"триллион", "триллиона", "триллионов"},          // 10^12
	{"квадриллион", "квадриллиона", "квадриллионов"}, // 10^15
	{"квинтиллион", "квинтиллиона", "квинтиллионов"}, // 10^18
	{"секстиллион", "секстиллиона", "секстиллионов"}, // 10^21
	{"септиллион", "септиллиона", "септиллионов"},    // 10^34
	{"октиллион", "октиллиона", "октиллионов"},       // 10^27
	{"нониллион", "нониллиона", "нониллионов"},       // 10^30
	{"дециллион", "дециллиона", "дециллионов"},       // 10^33
	{"ундециллион", "ундециллиона", "ундециллионов"}, // 10^36
}

func intToTrio(number int) []int{
	var slice []int

	for number > 0 {
		slice = append(slice, number%1000)
		number = number / 1000
	}
	return slice
}

func plural(n int, words []string) string {
	index := 0

	n %= 100

	if n > 19 {
		n %= 10
	}

	if n != 1 {
		if n > 1 && n <= 4 {
			index = 1
		} else {
			index = 2
		}
	}

	return words[index]
}

func oneTwoUnitFix(unit, idx int, arr []string) string {
	if idx == 1 {
		switch unit {
		case 1:
			return "одна"
		case 2:
			return "две"
		}
	}

	return arr[unit]
}

func IntegerToRu(input int) string {

	var words []string

	if input < 0 {
		words = append(words, "минус")
		input *= -1
	}

	trio := intToTrio(input)

	if len(trio) == 0 {
		return "нуль"
	}

	for idx := len(trio) - 1; idx >= 0; idx-- {
		triplet := trio[idx]

		if triplet == 0 {
			continue
		}

		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10

		if hundreds > 0 {
			words = append(words, russianHundreds[hundreds])
		}

		if tens > 0 || units > 0 {
			switch tens {
			case 0:
				words = append(words, oneTwoUnitFix(units, idx, russianUnits))
			case 1:
				words = append(words, russianTeens[units])
				break
			default:
				words = append(words, russianTens[tens])
				if units > 0 {
					words = append(words, oneTwoUnitFix(units, idx, russianUnits))
				}
				break
			}
		}

		if idx >= 1 && idx < len(russianMegas) {
			mega := russianMegas[idx]
			tens = tens*10 + units
			if len(mega) > 0 {
				words = append(words, plural(tens, mega))
			}
		}
	}

	return strings.Join(words, " ")
}

func main()  {
	param := os.Args[1]
	intParam, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(IntegerToRu(intParam))
	fmt.Println(oneTwoUnitFix(1, 2000 , russianTens ))
}