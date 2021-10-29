package main

import (
	"reflect"
	"testing"
)

func TestMakeSliceOfStrings(t *testing.T) {
	//Arrange
	testTable := []struct {
		values   		[]int
		expected 		string
	}{
		{
			values:  []int{4,4,1,6},
			expected: "4,4,1,6",
		},
		{
			values:  []int{1,456,111,666},
			expected: "1,456,111,666",
		},
	}

	//Act
	for _, testCase := range testTable {
		result := MakeSliceOfStrings(testCase.values)

		t.Logf("Calling (%d), result %s\n",
			testCase.values, testCase.expected)
		//Assert
		if reflect.DeepEqual(testCase.expected, result){
			t.Errorf("Incorrect result. Expect %s, got %s",
				testCase.expected, result)
		}
	}

}


func TestCountNums(t *testing.T) {
	//Arrange
	testTable := []struct {
		num    		float64
		expected 		[]int
	}{
		{
			num:  20,
			expected: []int{0,1,2,3},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := CountNums(testCase.num)

		t.Logf("Calling (%f), result %d\n",
			testCase.num, testCase.expected)
		//Assert
		if reflect.DeepEqual(testCase.expected, result){
			t.Errorf("Incorrect result. Expect %d, got %d",
				testCase.expected, result)
		}
	}

}