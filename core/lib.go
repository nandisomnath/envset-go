package core

import (
	"strings"
)

type ShellOptions int

const (
	Fish ShellOptions = iota
	Zsh
	Bash
	UnkownShell // this is just for handling returns
)

func GetShell(shell string) ShellOptions {

	if strings.Contains(shell, "zsh") {
		return Zsh
	} else if strings.Contains(shell, "bash") {
		return Bash
	} else if strings.Contains(shell, "fish") {
		return Fish
	} else {
		return UnkownShell
	}

}

func AddEnv(shell, envValue string, removePath bool) {
	shellOption := GetShell(shell)

	switch shellOption {
	case Zsh:
		panic("Zsh shell not implemented.")
	case Bash:
		panic("Bash shell not implemented.")
	case Fish:
		FishShellHandler(envValue, removePath)
	default:
		panic("Unkown shell not implemented.")
	}

}
