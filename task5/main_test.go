package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestParseFile(t *testing.T) {
	//Arrange
	testTable := []struct {
		input    		string
		expected 		string
	}{
		{
			input:  "126682",
			expected: "сто двадцать шесть тысяч шестьсот восемьдесят два",
		},
		{
			input:  "2935174315119",
			expected: "два триллиона девятьсот тридцать пять миллиардов сто семьдесят четыре миллиона триста пятнадцать тысяч сто девятнадцать",
		},
		{
			input:  "-1",
			expected: "минус один",
		},
	}

	//Act
	for _, testCase := range testTable {
		number, err := strconv.Atoi(testCase.input)
		if err != nil {
			fmt.Println(err)
		}
		result := IntegerToRu(number)

		t.Logf("Calling IntegerToRu(%d), result %v\n",
			number, testCase.expected)
		//Assert
		if result != testCase.expected{
			t.Errorf("Incorrect result. Expect %v, got %v",
				testCase.expected, result)
		}
	}
}