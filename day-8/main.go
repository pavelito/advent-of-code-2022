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
						// fmt.Println(fmt.Sprintf("%d %d is visible from top", x, y))
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
						// fmt.Println(fmt.Sprintf("%d %d is visible from left", x, y))
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
						// fmt.Println(fmt.Sprintf("%d %d is visible from bottom", x, y))
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
						// fmt.Println(fmt.Sprintf("%d %d is visible from right", x, y))
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

func getMaxScenicScore(trees [][]int) (maxScenicScore int) {

	maxScenicScore = 0
	score := 0
	leftReach, rightReach, topReach, bottomReach := 0, 0, 0, 0
	for x, row := range trees {
		for y, _ := range row {
			//outside trees
			if x == 0 ||
				x == (len(trees)-1) ||
				y == 0 ||
				y == (len(row)-1) {
				score = 0
			} else {
				//look right
				for r := y + 1; r <= len(row)-1; r++ {
					if r == len(row)-1 || (trees[x][r] >= trees[x][y]) {
						rightReach = r - y
						break
					}
				}
				//look bottom
				for b := x + 1; b <= len(trees)-1; b++ {
					if b == len(trees)-1 || (trees[b][y] >= trees[x][y]) {
						bottomReach = b - x
						break
					}
				}
				//look left
				for l := y - 1; l >= 0; l-- {
					if l == 0 || (trees[x][l] >= trees[x][y]) {
						leftReach = y - l
						break
					}
				}
				// //look top
				for t := x - 1; t >= 0; t-- {
					if t == 0 || (trees[t][y] >= trees[x][y]) {
						topReach = x - t
						break
					}
				}

			}
			//fmt.Println(x, y)
			//fmt.Println(topReach, leftReach, bottomReach, rightReach)
			score = leftReach * rightReach * topReach * bottomReach
			if score > maxScenicScore {
				maxScenicScore = score
				// fmt.Println(x, y, maxScenicScore)
			}
		}
	}

	return maxScenicScore
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
		var treeRow []int
		for _, v := range stringVal {
			treeHeight, _ := strconv.Atoi(string(v))
			treeRow = append(treeRow, treeHeight)
		}
		trees = append(trees, treeRow)
	}

	visibleTrees := getVisibleTrees(trees)
	// //fmt.Println(trees)
	// //fmt.Println(visibleTrees)
	fmt.Println("Part 1, Total Visible Trees: ", len(visibleTrees))
	fmt.Println("Part 2, Maximum Scenic Score: ", getMaxScenicScore(trees))

}
