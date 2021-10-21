
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func wordCount(path,searchedWord string)error{
	counts := map[string]int{}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		counts[word]++
	}
	if err != nil {
		return err
	}
	fmt.Println("Number of count is", counts[searchedWord])
	return nil
}



func replaceWord(word, newWord, path string){
	src := "modified.txt"
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(word), []byte(newWord), -1)
	if err = ioutil.WriteFile(src, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.Rename(src, path)
	if err != nil {
		return
	}
}


func main() {
	params := os.Args

	if len(params) == 3 {
		err := wordCount(params[1], params[2])
		if err != nil {
			return
		}
	} else if len(params) == 4 {
		replaceWord(params[1], params[2], params[3])
	}else {
		fmt.Println("To count words enter <searchedWord> <path>  or <word> <newWord> <path> if you want to replace some word \n")
	}
}