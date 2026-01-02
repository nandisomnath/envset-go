package core



// This definations is needed for a shell to work
type Shell interface {
	RemovePath()
	AddPath()
	InitSetup()
	GetShellConfigPath() string
	GetUserConfigPath() string
	GetShellCode() string
}