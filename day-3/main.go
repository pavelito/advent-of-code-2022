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
	for fileScanner.Scan() {
		stringVal := fileScanner.Text()

		sizeCompartment := len(stringVal) / 2
		compartment1 := stringVal[0:sizeCompartment]
		compartment2 := stringVal[sizeCompartment:]

		// fmt.Println("File Line", stringVal)
		// fmt.Println("Compartment 1 ", compartment1)
		// fmt.Println("Compartment 2 ", compartment2)
		commonLetter := findCommonLetter(compartment1, compartment2)
		priority := getPriority(commonLetter)
		fmt.Println("Common Letter is ", string(commonLetter))
		fmt.Println("Priority of common letter is ", priority)
		sumPriority += priority
	}
	fmt.Println("Total Priority ", sumPriority)
	readFile.Close()
}
