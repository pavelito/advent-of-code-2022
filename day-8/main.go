package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var trees [][]int

func getVisibleTrees(trees [][]int) [][2]int {
	var visibleTrees [][2]int

	for x, row := range trees {
		for y, _ := range row {
			//outside trees
			treeIsVisible := false
			if x == 0 ||
				x == (len(trees)-1) ||
				y == 0 ||
				y == (len(row)-1) {
				treeIsVisible = true
			} else {
				//look from top
				for t := 0; t <= x; t++ {
					if t == x {
						fmt.Println(fmt.Sprintf("%d %d is visible from top", x, y))
						treeIsVisible = true || treeIsVisible
						break
					}
					if trees[t][y] >= trees[x][y] {
						treeIsVisible = false || treeIsVisible //
						break
					}

				}
				//look from left
				for l := 0; l <= y; l++ {
					if l == y {
						fmt.Println(fmt.Sprintf("%d %d is visible from left", x, y))
						treeIsVisible = true || treeIsVisible
						break
					}
					if trees[x][l] >= trees[x][y] {
						treeIsVisible = false || treeIsVisible
						break
					}

				}
				//look from bottom
				for b := len(trees) - 1; b >= x; b-- {
					if b == x {
						fmt.Println(fmt.Sprintf("%d %d is visible from bottom", x, y))
						treeIsVisible = true || treeIsVisible
						break
					}
					if trees[b][y] >= trees[x][y] {
						treeIsVisible = false || treeIsVisible
						break
					}

				}
				//look from right
				for r := len(row) - 1; r >= y; r-- {
					if r == y {
						fmt.Println(fmt.Sprintf("%d %d is visible from row", x, y))
						treeIsVisible = true || treeIsVisible
						break
					}
					if trees[x][r] >= trees[x][y] {
						treeIsVisible = false || treeIsVisible
						break
					}

				}

			}
			if treeIsVisible {
				visibleTrees = append(visibleTrees, [2]int{x, y})
			}

		}
	}
	return visibleTrees
}

func main() {

	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		stringVal := fileScanner.Text()
		//fmt.Println("File Line", stringVal)
		var treeRow []int
		for _, v := range stringVal {
			treeHeight, _ := strconv.Atoi(string(v))
			treeRow = append(treeRow, treeHeight)
		}
		trees = append(trees, treeRow)
	}

	visibleTrees := getVisibleTrees(trees)
	//fmt.Println(trees)
	//fmt.Println(visibleTrees)
	fmt.Println(len(visibleTrees))

}
