package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	var maxCal = 0
	var total = 0
	var maxElfCount = 0
	var elfCount = 0
	var elfCalSlice []int

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		stringVal := fileScanner.Text()
		intValue, err := strconv.Atoi(stringVal)

		if err == nil {
			total = total + intValue
		} else {
			elfCount = elfCount + 1
			if total > maxCal {
				maxCal = total
				maxElfCount = elfCount
			}

			elfCalSlice = append(elfCalSlice, total)
			// fmt.Println("Processed Elf ", elfCount)
			// fmt.Println("Elf Total", total)
			total = 0
		}
	}

	readFile.Close()
	fmt.Println("The max calories is ", maxCal)
	fmt.Println("The Elf who has the max calories is ", maxElfCount)
	sort.Sort(sort.Reverse(sort.IntSlice(elfCalSlice)))
	fmt.Println("The Top 3 Elf calories slice ", elfCalSlice[0:3])
	sumTop3 := 0
	for _, v := range elfCalSlice[0:3] {
		sumTop3 += v
	}
	fmt.Println("The Sum of Top 3 calories is ", sumTop3)

}
