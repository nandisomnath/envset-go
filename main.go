package main

import (
	"github.com/nandisomnath/envset/core"
	"github.com/spf13/pflag"
)

func main() {
	removePath := pflag.BoolP("remove","r", false, "remove the folder from the path")
	pathValue := pflag.StringP("path", "p","", "folder path to add in $PATH variable")
	shellValue := pflag.StringP("shell", "s","", "which shell the path will be added")
	pflag.Parse()
	core.AddEnv(*shellValue, *pathValue, *removePath)
}
