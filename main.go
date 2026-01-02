package main

import (
	"fmt"
	"os"
	"github.com/nandisomnath/envset/core"
	"github.com/spf13/pflag"
)

func main() {
	removePath := pflag.BoolP("remove", "r", false, "remove the folder from the path")
	pathValue := pflag.StringP("path", "p", "", "folder path to add in $PATH variable")
	shellValue := pflag.StringP("shell", "s", "", "which shell the path will be added")
	version := pflag.BoolP("version", "v", false, "Shows version and exit")
	maintainers := pflag.BoolP("maintainers", "m", false, "Shows maintainers and exit")
	pflag.Parse()

	if *version && !*maintainers {
		fmt.Printf("%s", App.InfoString())
		os.Exit(0)
	}

	if *maintainers && !*version {
		for _, m := range App.Maintainers {
			fmt.Println(m)
		}
		os.Exit(0)
	}

	core.AddEnv(*shellValue, *pathValue, *removePath)
}
