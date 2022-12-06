package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkUnique(array []string) bool {
	for k, v := range array {
		for _, w := range array[k+1:] {
			if v == w {
				return false
			}
		}
	}
	return true
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		stringVal := fileScanner.Text()
		distinctCharCount := 14
		for i := distinctCharCount - 1; i < len(stringVal); i++ {
			var bufferArray []string
			for k := i - (distinctCharCount - 1); k <= i; k++ {
				bufferArray = append(bufferArray, string(stringVal[k]))
			}
			if checkUnique(bufferArray) {
				fmt.Println("Uniqe key is ", i+1)
				break
			}
		}

	}

	readFile.Close()
}
