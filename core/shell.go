package core

import "path"


type Shell struct {
	// name of the shell we are working with
	Name string 
	// config path where the main config file of that shell exists
	ConfigPath string
}


func NewShell(name, _path string) Shell {
	return Shell {
		ConfigPath: path.Clean(_path),
		Name: name,
	}
}


