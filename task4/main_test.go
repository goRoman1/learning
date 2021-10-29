package main

import "testing"

func TestWordCount(t *testing.T) {
	//Arrange
	testTable := []struct {
		path    		string
		searchedWord 	string
		expected 		int
	}{
		{
			path:  "text.txt",
			searchedWord :  "ok",
			expected: 3,
		},
		{
			path:  "text.txt",
			searchedWord :  "and",
			expected: 2,
		},
	}

	//Act
	for _, testCase := range testTable {
		result := WordCount(testCase.path, testCase.searchedWord)

		t.Logf("Calling TriangleArea(%s%s), result %d\n",
			testCase.path, testCase.searchedWord, testCase.expected)
		//Assert
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %d, got %d",
				testCase.expected, result)
		}
	}

}

/*
func TestReplaceWord(t *testing.T) {
	//Arrange
	testTable := []struct {
		path    		string
		searchedWord 	string
		expected 		int
	}{
		{
			path:  "text.txt",
			searchedWord :  "ok",
			expected: 3,
		},
		{
			path:  "text.txt",
			searchedWord :  "and",
			expected: 2,
		},
	}

	//Act
	for _, testCase := range testTable {
		result := WordCount(testCase.path, testCase.searchedWord)

		t.Logf("Calling TriangleArea(%s%s), result %d\n",
			testCase.path, testCase.searchedWord, testCase.expected)
		//Assert
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %d, got %d",
				testCase.expected, result)
		}
	}

}
*/