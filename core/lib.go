package core

import (
	"os"
	"strings"
)






func GetShell() Shell {
	
	shell := os.Getenv("SHELL")

	if strings.Contains(shell, "zsh") {
		return ZSH
	} else if strings.Contains(shell, "bash") {
		return BASH
	} else if strings.Contains(shell, "fish") {
		return FISH
	} else {
		return UNKOWN_SHELL
	}

}



func AddEnv(envName, envValue string)  {
	
}

