package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type section struct {
	start int
	end   int
}

func fullyContains(a section, b section) bool {

	if a.start >= b.start && a.end <= b.end {
		fmt.Println(a, " is inside ", b)
		return true
	} else if b.start >= a.start && b.end <= a.end {
		fmt.Println(b, " is inside ", a)
		return true
	}
	return false
}

func overalps(a section, b section) bool {
	if b.start <= a.end || a.start <= b.end {
		return true
	}
	return false
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	fullyContainsRanges := 0

	for fileScanner.Scan() {
		stringVal := fileScanner.Text()
		sections := strings.Split(stringVal, ",")
		secA := strings.Split(sections[0], "-")
		secB := strings.Split(sections[1], "-")

		parsedIntStart, _ := strconv.Atoi(secA[0])
		parsedIntEnd, _ := strconv.Atoi(secA[1])
		sectionA := section{start: parsedIntStart, end: parsedIntEnd}

		parsedIntStart, _ = strconv.Atoi(secB[0])
		parsedIntEnd, _ = strconv.Atoi(secB[1])
		sectionB := section{start: parsedIntStart, end: parsedIntEnd}

		fmt.Println("File Line", stringVal)
		fmt.Println("Section A ", sectionA)
		fmt.Println("Section B ", sectionB)
		if fullyContains(sectionA, sectionB) {
			fullyContainsRanges++
		}
	}

	readFile.Close()
	fmt.Println("No of fully contains - ", fullyContainsRanges)
}
