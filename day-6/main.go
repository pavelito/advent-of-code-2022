package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkUnique(array [14]string) bool {
	for k, v := range array {
		for _, w := range array[k+1:] {
			if v == w {
				//fmt.Println(v, " is same as ", w)
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
		// var bufferArray [4]string
		// stringVal := fileScanner.Text()
		// //fmt.Println("File Line", stringVal)
		// for i := 3; i < len(stringVal); i++ {
		// 	bufferArray[0] = string(stringVal[i-3])
		// 	bufferArray[1] = string(stringVal[i-2])
		// 	bufferArray[2] = string(stringVal[i-1])
		// 	bufferArray[3] = string(stringVal[i])
		// 	//fmt.Println(bufferArray)
		// 	if checkUnique(bufferArray) {
		// 		fmt.Println("Uniqe key is ", i+1)
		// 		break
		// 	}
		// }

		var bufferArray [14]string
		stringVal := fileScanner.Text()
		//fmt.Println("File Line", stringVal)
		for i := 13; i < len(stringVal); i++ {
			j := 0
			for k := i - 13; k <= i; k++ {
				bufferArray[j] = string(stringVal[k])
				j++
			}
			//fmt.Println(bufferArray)
			if checkUnique(bufferArray) {
				fmt.Println("Uniqe key is ", i+1)
				break
			}
		}

	}

	readFile.Close()
}
