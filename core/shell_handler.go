package core

import (
	"encoding/json"
	"fmt"
)






// this is the wrapper for shell handle
type ShellHandler struct {
	Shells []Shell
}


// Creates New shell handler which retrives rules for different shells
func NewShellHandler() (v *ShellHandler, err error)  {
	shell_handler := ShellHandler {
		Shells: make([]Shell, 0),
	}

	data := []byte(RULES)
	err = json.Unmarshal(data, &shell_handler.s)
	if err != nil {
		return nil, err
	}

	fmt.Println(shell_handler.Shells[0])

	return &shell_handler, nil
}

