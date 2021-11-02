package main

import (
	"fmt"
//	"reflect"
	"strconv"
	"testing"
)

func TestIntegerToRu(t *testing.T) {
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
/*
func TestIntToTrio(t *testing.T) {
	//Arrange
	testTable := []struct {
		number    	int
		expected 	[]int
	}{
		{
			number:  1234,
			expected: []int{234,1},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := intToTrio(testCase.number)

		t.Logf("Calling intToTrio(%d), result %v\n",
			testCase.number, testCase.expected)
		//Assert
		if reflect.DeepEqual(result, testCase.expected){
			t.Errorf("Incorrect result. Expect %v, got %v",
				testCase.expected, result)
		}
	}
}


func TestOneTwoUnitFix(t *testing.T) {
	//Arrange
	testTable := []struct {
		unit, idx    	int
		arr             []string
		expected 		string
	}{
		{
			unit:  1,
			idx : 2,
			arr: []string{"",""},
			expected: "одна",
		},
	}

	//Act
	for _, testCase := range testTable {
		result := oneTwoUnitFix(testCase.unit,testCase.idx,testCase.arr)

		t.Logf("Calling IntegerToRu(%d, %d,%v), result %s\n",
			testCase.unit, testCase.idx, testCase.arr, testCase.expected)
		//Assert
		if result != testCase.expected{
			t.Errorf("Incorrect result. Expect %v, got %v",
				testCase.expected, result)
		}
	}
}

func TestPlural(t *testing.T) {
	//Arrange
	testTable := []struct {
		n    		int
		words    	[]string
		expected 	string
	}{
		{
			n: 5,
			words: []string{"",""},
			expected: "одна",
		},
	}

	//Act
	for _, testCase := range testTable {
		result := plural(testCase.n,testCase.words)

		t.Logf("Calling IntegerToRu(%d,%v), result %s\n",
			testCase.n, testCase.words, testCase.expected)
		//Assert
		if result != testCase.expected{
			t.Errorf("Incorrect result. Expect %v, got %v",
				testCase.expected, result)
		}
	}
}
 */