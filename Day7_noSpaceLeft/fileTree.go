package main

import (
	"math"
	"sort"
)

type FileTree struct {
	value  File
	childs map[string]*FileTree
}

func NewFileTree() FileTree {
	return FileTree{childs: make(map[string]*FileTree, 0)}
}

func NewFileTreeFromFileSystem(fileSystem FileSystem) (fileTree FileTree) {
	fileTree = NewFileTree()
	for _, file := range fileSystem.files {
		fileTree.File(file)
	}
	return
}

func (fileTree *FileTree) File(file File) {
	var fileFileTree *FileTree = fileTree
	for idx, directory := range file.path.path {
		if idx == 0 {
			continue
		}
		if _, ok := fileFileTree.childs[directory]; !ok {
			var newFileTree = NewFileTree()
			fileFileTree.childs[directory] = &newFileTree
		}
		fileFileTree = fileFileTree.childs[directory]
	}
	fileFileTree.childs[file.name] = &FileTree{file, make(map[string]*FileTree, 0)}
}

func (fileTree *FileTree) Evaluate() {
	var size int
	for _, currentFileTree := range fileTree.childs {
		if currentFileTree.value.name == "" {
			currentFileTree.Evaluate()
		}
		size += currentFileTree.value.size
	}
	fileTree.value = NewFile("", size, Path{})
}

func (fileTree *FileTree) CountDirectory(function func(file File) bool) (sum int) {
	for _, child := range fileTree.childs {
		if child.value.name != "" {
			continue
		}
		if function(child.value) {
			sum += child.value.size
		}
		sum += child.CountDirectory(function)
	}
	return
}

func (fileTree *FileTree) SmallestDirectoryToDelete(maxCapacity int, neededSpace int) int {
	var freeSpace = maxCapacity - fileTree.value.size
	var deleteSize = neededSpace - freeSpace
	var directories = fileTree.GetAllDirectories()
	for i := 0; i < len(directories); i++ {
		if directories[i].value.size < deleteSize {
			directories = SliceRemove(directories, i)
			i--
			continue
		}
	}
	sort.Slice(directories, func(a int, b int) bool {
		return math.Abs(float64(deleteSize-directories[a].value.size)) < math.Abs(float64(deleteSize-directories[b].value.size))
	})
	return directories[0].value.size
}

func (fileTree *FileTree) GetAllDirectories() (directories []*FileTree) {
	directories = make([]*FileTree, 0)
	for _, child := range fileTree.childs {
		if child.value.name == "" {
			directories = append(directories, child)
			directories = append(directories, child.GetAllDirectories()...)
		}
	}
	return
}
