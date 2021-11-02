package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestWordCount(t *testing.T) {
	//Arrange
	testTable := []struct {
		path    		string
		searchedWord 	string
		expected 		int
	}{
		{
			path:  "test1.txt",
			searchedWord :"santa",
			expected: 3,
		},
	}

	//Act
	for _, testCase := range testTable {
		result := WordCount(testCase.path, testCase.searchedWord)

		t.Logf("Calling TestWordCount(%s%s), result %d\n",
			testCase.path, testCase.searchedWord, testCase.expected)
		//Assert
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %d, got %d",
				testCase.expected, result)
		}
	}

}

func TestReplaceWord(t *testing.T) {
	//Arrange
	testTable := []struct {
		path    		string
		word		 	string
		newWord			string
		expected 		bool
	}{
		{
			path:  "test2.txt",
			word: "jfd",
			newWord: "ok",
			expected: true,
		},
	}

	//Act
	for _, testCase := range testTable {
		input, err := ioutil.ReadFile(testCase.path)
		if err != nil {
			fmt.Println(err)
			return
		}

		output := ReplaceWord(testCase.path, testCase.word, testCase.newWord)
		compare := bytes.Compare(input,output)

		err = ioutil.WriteFile(testCase.path, input, 0644)
		if err != nil {
			log.Fatal(err)
		}


		//Assert
		if  compare == 0 && output != nil {
			t.Errorf("The content of the files has not been changed")
		} else if output == nil {
			t.Errorf("The error occured during function execution")
		} else {
			t.Logf("Calling ReplaceWord(%s%s%s), \n",
				testCase.path, testCase.word, testCase.newWord)
		}
	}

}

