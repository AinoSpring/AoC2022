package main

import (
	"fmt"
)

func main() {
	var session = Import()
	fmt.Println()
	var fileTree = NewFileTreeFromFileSystem(session.fileSystem)
	fileTree.Evaluate()
	fmt.Println(fileTree.CountDirectory(func(file File) bool {
		return file.size <= 100000
	}))
	fmt.Print("\n------Part two------\n\n")
	fmt.Println(fileTree.SmallestDirectoryToDelete(70000000, 30000000))
	fmt.Println()
}
