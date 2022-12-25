package main

type Path struct {
	path []string
}

func NewPath() Path {
	return Path{[]string{"/"}}
}

func (path *Path) changeDirectory(value string) {
	if value == ".." {
		path.path = path.path[:(len(path.path) - 1)]
		return
	}
	if value == "/" {
		path.path = []string{"/"}
		return
	}
	path.path = append(path.path, value)
}

func (path *Path) Copy() Path {
	var newPath = make([]string, len(path.path))
	copy(newPath, path.path)
	return Path{newPath}
}

func (path *Path) Directory() string {
	return path.path[len(path.path)-1]
}
