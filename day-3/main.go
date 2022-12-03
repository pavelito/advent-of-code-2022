package main

import (
	"bufio"
	"fmt"
	"os"
)

func findCommonLetter(a string, b string) rune {
	for _, charA := range a {
		for _, charB := range b {
			if charA == charB {
				//return string(charA)
				return charA
			}
		}
	}
	return 0
}

func findCommonLetterPart2(a string, b string, c string) rune {
	var commonLettersFirstPass []rune
	for _, charA := range a {
		for _, charB := range b {
			if charA == charB {
				commonLettersFirstPass = append(commonLettersFirstPass, charA)
			}
		}
	}
	for _, charCommon := range commonLettersFirstPass {
		for _, charC := range c {
			if charCommon == charC {
				return charCommon
			}
		}
	}
	return 0
}

func getPriority(commonLetter rune) int {
	if commonLetter >= 65 && commonLetter < 97 {
		return int(commonLetter - 38)
	}
	return int(commonLetter - 96)

}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	sumPriority := 0
	lineCounter := 0
	var groupOfBags [3]string
	for fileScanner.Scan() {
		stringVal := fileScanner.Text()
		moduloLine := lineCounter % 3
		//Part 2
		groupOfBags[moduloLine] = stringVal

		if moduloLine == 2 {
			//calculate group
			commonLetter := findCommonLetterPart2(groupOfBags[0], groupOfBags[1], groupOfBags[2])
			priority := getPriority(commonLetter)
			fmt.Println("Common Letter is ", string(commonLetter))
			fmt.Println("Priority of common letter is ", priority)
			sumPriority += priority
		}
		lineCounter++

		//Part 1
		// sizeCompartment := len(stringVal) / 2
		// compartment1 := stringVal[0:sizeCompartment]
		// compartment2 := stringVal[sizeCompartment:]

		// // fmt.Println("File Line", stringVal)
		// // fmt.Println("Compartment 1 ", compartment1)
		// // fmt.Println("Compartment 2 ", compartment2)
		// commonLetter := findCommonLetter(compartment1, compartment2)
		// priority := getPriority(commonLetter)
		// fmt.Println("Common Letter is ", string(commonLetter))
		// fmt.Println("Priority of common letter is ", priority)
		// sumPriority += priority
	}
	fmt.Println("Total Priority ", sumPriority)
	readFile.Close()
}