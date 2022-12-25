package main

type FileSystem struct {
	files []File
}

func NewFileSystem() FileSystem {
	return FileSystem{make([]File, 0)}
}

func (fileSystem *FileSystem) File(file File) {
	fileSystem.files = append(fileSystem.files, file)
}
