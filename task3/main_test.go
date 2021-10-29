package main

import (
	"math"
	"testing"
)
/*
func TestContinueConfirmation(t *testing.T) {
	testTable := []struct{
		ask  string
		expected bool
	}{
		{
			ask: "yes",
			expected: true,
		},
		{
			ask: "y",
			expected: true,
		},
	}

	for _, testCase := range testTable {
		result := continueConfirmation()
		t.Logf("Calling CanBeTriangle(%f,%f,%f), result %v\n",
			testCase.a,testCase.b,testCase.c, testTable)


		if result != testCase.expected {
			t.Errorf("Incorrect expected. Expect expected %v but got %v -> for input params num = %v, %v, %v", testCase.expected, result, testCase.a, testCase.b, testCase.c)
		}
	}
}
 */

func TestCanBeTriangle(t *testing.T) {
	testTable := []struct{
		a, b, c  float64
		expected bool
	}{
		{
			a: 2,
			b: 3,
			c: 4,
			expected: true,
		},
		{
			a: 1,
			b: 2,
			c: 3,
			expected: false,
		},
	}

	for _, testCase := range testTable {
		result := IsTriangle(testCase.a, testCase.b, testCase.c)
		t.Logf("Calling CanBeTriangle(%f,%f,%f), result %v\n",
			testCase.a,testCase.b,testCase.c, testTable)


		if result != testCase.expected {
			t.Errorf("Incorrect expected. Expect expected %v but got %v -> for input params num = %v, %v, %v", testCase.expected, result, testCase.a, testCase.b, testCase.c)
		}
	}
}

func TestParseTriangleInput(t *testing.T) {
	testTable := []struct {
		input                           string
		expectedA, expectedB, expectedC float64
		expectedName                    string
	}{
		{
			input : "tr1,2,3,4",
			expectedA: 2.0,
			expectedB:3.0,
			expectedC:4.0,
			expectedName: "tr1" ,
		},
		{
			input : "2,3,4",
			expectedA:0.0,
			expectedB:0.0,
			expectedC:0.0,
			expectedName: "",
		},
		{
			input : "tr1,2,3,ab",
			expectedA: 0.0,
			expectedB:0.0,
			expectedC:0.0,
			expectedName: "" ,
		},
	}

	for _, testCase := range testTable {
		resultName, resultA, resultB, resultC := ParseInputStr(testCase.input)
		t.Logf("Calling ParseTriangleInput(%s), result %s, %f,%f,%f \n",
			testCase.input, testCase.expectedName, testCase.expectedA,testCase.expectedB,testCase.expectedC)

		if resultA != testCase.expectedA || resultB != testCase.expectedB || resultC != testCase.expectedC || resultName != testCase.expectedName {
			t.Errorf("parsing error on params %v", testCase)
		}
	}
}

func TestTrimTabsAndSpaces(t *testing.T) {
	testTable := [][2]string{
		{" ", ""},
		{"\t", ""},
		{"\t  \t", ""},
		{" \tabc \t", "abc"},
	}

	for _, testCase := range testTable {
		result := TrimTabsAndSpaces(testCase[0])
		t.Logf("Calling TrimTabsAndSpaces(%v), result %v\n",
			testCase[0], testTable)

		if result != testCase[1] {
			t.Errorf("expected <%v> but received <%v> -> for param %v", testCase[1], result, testCase[0])
		}
	}
}

func TestTriangleArea(t *testing.T) {
	testTable := []struct {
		a,b,c    		float64
		expected 		float64
	}{
		{
			a: 4,
			b: 5,
			c: 6,
			expected: 9.92156,
		},
		{
			a: 1,
			b: 1,
			c: 1,
			expected: 0.43301270,
		},
	}

	for _, testCase := range testTable {
		result := TriangleArea(testCase.a, testCase.b, testCase.c)
		t.Logf("Calling TriangleArea(%f%f%f), result %v\n",
			testCase.a,testCase.b, testCase.c, testTable)


		if math.Abs(result - testCase.expected) >= 0.001 {
			t.Errorf("expected %v but received %v -> for triangle %f,%f,%f,", testCase.expected, result, testCase.a, testCase.b, testCase.c)
		}
	}
}


