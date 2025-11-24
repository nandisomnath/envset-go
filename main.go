package main

import (
	"fmt"

	"github.com/nandisomnath/envset/core"
)


func main() {
	shell_handler, err := core.InitHandler();

	if err != nil {
		panic(err)
	}

	fmt.Println(shell_handler.S())

}



