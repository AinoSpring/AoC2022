package main

type File struct {
	name string
	size int
	path Path
}

func NewFile(name string, size int, path Path) File {
	return File{name, size, path}
}
