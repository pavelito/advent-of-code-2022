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

//The trick here is, store ALL the common letters between a and b
//Whereas the first common letter between a+b and c (thus reusing the function for stage 2)
func findCommonLetterPart2(a string, b string, c string) rune {
	var commonLettersFirstPass []rune
	for _, charA := range a {
		for _, charB := range b {
			if charA == charB {
				commonLettersFirstPass = append(commonLettersFirstPass, charA)
			}
		}
	}
	return findCommonLetter(string(commonLettersFirstPass), c)
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
	sumPriorityPart1 := 0
	sumPriorityPart2 := 0
	lineCounter := 0
	var groupOfBags [3]string
	for fileScanner.Scan() {
		stringVal := fileScanner.Text()

		//Part 1
		//8394
		sizeCompartment := len(stringVal) / 2
		compartment1 := stringVal[0:sizeCompartment]
		compartment2 := stringVal[sizeCompartment:]

		commonLetter := findCommonLetter(compartment1, compartment2)
		priority := getPriority(commonLetter)

		sumPriorityPart1 += priority

		//Part 2
		//2413
		moduloLine := lineCounter % 3
		groupOfBags[moduloLine] = stringVal

		if moduloLine == 2 {
			//calculate group
			commonLetter := findCommonLetterPart2(groupOfBags[0], groupOfBags[1], groupOfBags[2])
			priority := getPriority(commonLetter)

			sumPriorityPart2 += priority
		}
		lineCounter++
	}
	fmt.Println("Total Priority Part 1", sumPriorityPart1)
	fmt.Println("Total Priority Part 2", sumPriorityPart2)
	readFile.Close()
}
