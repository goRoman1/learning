package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReplaceWord(path, word, newWord string)[]byte{
	src := "modified.txt"
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	output := bytes.Replace(input, []byte(word), []byte(newWord), -1)
	fmt.Println(output)
	if err = ioutil.WriteFile(src, output, 0666); err != nil {
		fmt.Println(err)
		return nil
	}

	err = os.Rename(src, path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return output
}

func WordCount(path,searchedWord string)int{
	counts := map[string]int{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)  //
	scanner.Split(bufio.ScanWords)     //
	for scanner.Scan() {               //
		word := scanner.Text()         //
		word = strings.ToLower(word)   //
		counts[word]++                 //
	}
	if err != nil {
		fmt.Println(err)
	}

//	fmt.Println("Number of count is", counts[searchedWord])
	return counts[searchedWord]
}


func main() {
	params := os.Args

	if len(params) == 3 {
		_, err := fmt.Println(WordCount(params[1], params[2]))
		if err != nil {
			return
		}
		return
	} else if len(params) == 4 {
		ReplaceWord(params[1], params[2], params[3])
		return
	}

	fmt.Println("To count words enter <path> <searchedWord>  or <path> <word> <newWord>  if you want to replace some word")
}