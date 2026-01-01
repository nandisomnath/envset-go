package core

import (
	"strings"
)


type ShellOptions int

const (
	Fish ShellOptions = iota 
	Zsh
	Bash
	UnkownShell 
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



func AddEnv(envName, envValue string)  {
	
}

