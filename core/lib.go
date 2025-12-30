package core

import (
	"os"
	"strings"
)


type ShellOptions int
const (
	Fish ShellOptions = iota 
	Zsh
	Bash
	UnkownShell 
)



func GetShell(shell: string) ShellOptions {

	if strings.Contains(shell, "zsh") {
		return Zsh 
	} else if strings.Contains(shell, "bash") {
		return Bash 
	} else if strings.Contains(shell, "fish") {
		return Fish 
	} else {
		return 
	}

}



func AddEnv(envName, envValue string)  {
	
}

