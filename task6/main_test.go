package main

import (
	"reflect"
	"testing"
)

func TestToIntSlice(t *testing.T) {
	//Arrange
	testTable := []struct {
		str    		string
		expected 	[]int
	}{
		{
			str:  "143211",
			expected: []int{1,4,3,2,1,1},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := toIntSlice(testCase.str)

		t.Logf("Calling ToIntSlice(%s), result %d\n",
			testCase.str, testCase.expected)
		//Assert
		if !reflect.DeepEqual(testCase.expected, result){
			t.Errorf("Incorrect result. Expect %d, got %d",
				testCase.expected, result)
		}
	}

}

func TestIsPiter(t *testing.T) {
	//Arrange
	testTable := []struct {
		input    	[]int
		expected 	bool
	}{
		{
			input: []int{1,4,3,2,1,1},
			expected: false,
		},
	}

	//Act
	for _, testCase := range testTable {
		result := isPiter(testCase.input)

		t.Logf("Calling isPiter(%d), result %v\n",
			testCase.input, testCase.expected)
		//Assert
		if len(testCase.input) != 6 {
			t.Errorf("Incorrect result. Expect input len  = 6, got %d ",
				len(testCase.input))
		}

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v",
				testCase.expected, result)
		}

	}
}

func TestIsMoskow(t *testing.T) {
	//Arrange
	testTable := []struct {
		input    	[]int
		expected 	bool
	}{
		{
			input: []int{1,4,3,2,1,1},
			expected: false,
		},
	}

	//Act
	for _, testCase := range testTable {
		result := isMoskow(testCase.input)

		t.Logf("Calling isMoscow(%d), result %v\n",
			testCase.input, testCase.expected)
		//Assert
		if len(testCase.input) != 6 {
			t.Errorf("Incorrect result. Expect input len  = 6, got %d ",
				len(testCase.input))
		}

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v",
				testCase.expected, result)
		}

	}
}

func TestParseFile(t *testing.T) {
	//Arrange
	testTable := []struct {
		path    		string
		expected 		[]string
	}{
		{
			path:  "test.txt",
			expected: []string{"piter","131113", "333222", "641146", "325910"},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := parseFile(testCase.path)

		t.Logf("Calling ParseFile(%s), result %v\n",
			testCase.path, testCase.expected)
		//Assert
		if !reflect.DeepEqual(result, testCase.expected){
			t.Errorf("Incorrect result. Expect %v, got %v",
				testCase.expected, result)
		}
	}

}
