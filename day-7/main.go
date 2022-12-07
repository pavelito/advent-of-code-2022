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
	//fmt.Println(fmt.Sprintf("%s is of size %d and has parent %p", dir.path, dir.fileSize, dir.parent))
	if dir.fileSize <= 100000 {
		fmt.Println(fmt.Sprintf("%s is of size %d and has parent %p", dir.path, dir.fileSize, dir.parent))
		*total += dir.fileSize
	}
	if dir.children == nil {
		return
	}
	for _, v := range dir.children {
		traverseDirectory(v, total)
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
			//fmt.Println(consoleOut)
			fileSize, _ := strconv.Atoi(consoleOut[0])
			// currentDirectory.fileSize += fileSize

			//now add the filesize to current dir and all the way up to root parent
			incrementFileSize(currentDirectory, fileSize)

		}

		//fmt.Println(currentDirectory)
	}
	//fmt.Println(directoryTree)
	// fmt.Println(directoryTree.children["a"])

	total := 0
	traverseDirectory(&directoryTree, &total)
	fmt.Println(total)
}
