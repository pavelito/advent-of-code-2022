package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	path     string
	fileSize int
	children map[string]*Directory
	parent   *Directory
}

func incrementFileSize(dir *Directory, size int) {
	dir.fileSize += size
	if dir.parent == nil {
		return
	}
	incrementFileSize(dir.parent, size)
}

func traverseDirectory(dir *Directory, total *int) {
	if dir.fileSize <= 100000 {
		*total += dir.fileSize
	}
	if dir.children == nil {
		return
	}
	for _, v := range dir.children {
		traverseDirectory(v, total)
	}
}

func findMinimumSize(dir *Directory, targetSize int, resultSize *int) {
	//fmt.Println(fmt.Sprintf("%s is of size %d and has parent %p", dir.path, dir.fileSize, dir.parent))
	if dir.fileSize >= targetSize {
		//fmt.Println(fmt.Sprintf("%s is of size %d and has parent %p", dir.path, dir.fileSize, dir.parent))
		if dir.fileSize < *resultSize {
			//fmt.Println("Updating result size to ", dir.fileSize)
			*resultSize = dir.fileSize
		}

	}
	if dir.children == nil {
		return
	}
	for _, v := range dir.children {
		findMinimumSize(v, targetSize, resultSize)
	}
}

func main() {

	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	directoryTree := Directory{
		path: "/",
	}
	var currentDirectory *Directory
	for fileScanner.Scan() {

		stringVal := fileScanner.Text()
		consoleOut := strings.Fields(stringVal)
		if consoleOut[0] == "$" {
			if consoleOut[1] == "cd" {
				if consoleOut[2] == "/" {
					currentDirectory = &directoryTree
				} else if consoleOut[2] == ".." {
					currentDirectory = currentDirectory.parent
				} else {
					goToChildrenDirectory := currentDirectory.children[consoleOut[2]]
					currentDirectory = goToChildrenDirectory
				}
			} else if consoleOut[0] == "ls" {

			}
		} else if consoleOut[0] == "dir" {
			if currentDirectory.children == nil {
				currentDirectory.children = make(map[string]*Directory)
			}
			currentDirectory.children[consoleOut[1]] = &Directory{
				path:     consoleOut[1],
				fileSize: 0,
				parent:   currentDirectory,
			}
		} else { //fileout
			fileSize, _ := strconv.Atoi(consoleOut[0])

			//now add the filesize to current dir and all the way up to root parent
			incrementFileSize(currentDirectory, fileSize)

		}
	}

	//Part 1
	total := 0
	traverseDirectory(&directoryTree, &total)
	fmt.Println("Part 1 - Sum of all directories that are at most (max) 100000", total)

	//Part 2
	totalDiskSize := 70000000
	updateDiskSize := 30000000
	minimumCleanupSize := updateDiskSize - (totalDiskSize - directoryTree.fileSize)
	fmt.Println("Root has size ", directoryTree.fileSize)
	fmt.Println("Total Current Free size ", totalDiskSize-directoryTree.fileSize)
	fmt.Println("Minimum Required Free size ", minimumCleanupSize)

	resultSize := totalDiskSize
	findMinimumSize(&directoryTree, minimumCleanupSize, &resultSize)
	fmt.Println("Part 2 - Minimum Available Free size to Cleanup ", resultSize)
}
