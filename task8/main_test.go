package main

import (
	"reflect"
	"testing"
)

func TestCountFibs(t *testing.T) {
	//Arrange
	testTable := []struct {
		low,high    	int
		expected 		[]int
	}{
		{
			low: 3 ,
			high: 22  ,
			expected: []int{ 5, 8, 13, 21},
		},
		{
			low: 9 ,
			high: 60  ,
			expected: []int{ 13, 21, 34, 55},
		},
	}
	//Act
	for _, testCase := range testTable {
		result := countFibs(testCase.low, testCase.high)

		t.Logf("Calling TriangleArea(%d%d), result %d\n",
			testCase.low, testCase.high, testCase.expected)
		//Assert
		if reflect.DeepEqual(result, testCase.expected){
			t.Errorf("Incorrect result. Expect %d, got %d",
				testCase.expected, result)
		}
	}
}

func TestMakeString(t *testing.T) {
	//Arrange
	testTable := []struct {
		intSlice    	[]int
		expected 		string
	}{
		{
			intSlice: []int{1,2,3,4},
			expected: "1,2,3,4" ,
		},
		{
			intSlice: []int{1,6,3,1},
			expected: "1,6,3,1" ,
		},
	}
	//Act
	for _, testCase := range testTable {
		result := makeString(testCase.intSlice)

		t.Logf("Calling TriangleArea(%d), result %s\n",
			testCase.intSlice, testCase.expected)
		//Assert
		if reflect.DeepEqual(result, testCase.expected){
			t.Errorf("Incorrect result. Expect %s, got %s",
				testCase.expected, result)
		}
	}
}
