package core

import "os"

// this is the wrapper for shell handle
type ShellHandler struct {
	CurrentShell Shell
}



func NewShellHandler(shell Shell) ShellHandler {
	return ShellHandler {
		CurrentShell: shell,
	}
}


func WriteEnv(shell Shell, data string) error {
	fileName := shell.ConfigPath	
	return os.WriteFile(fileName, []byte(data), os.ModeAppend)
}

