package main

type Session struct {
	fileSystem              FileSystem
	currentWorkingDirectory Path
}

func NewSession() Session {
	return Session{fileSystem: NewFileSystem(), currentWorkingDirectory: NewPath()}
}

func (session *Session) NewFile(name string, size int) {
	session.fileSystem.File(NewFile(name, size, session.currentWorkingDirectory.Copy()))
}
