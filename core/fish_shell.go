package core

import (
	"fmt"
	"os"
	"path"
)

type Shell interface {
	RemovePath()
	AddPath()
	InitSetup()
	GetShellConfigPath() string
	GetUserConfigPath() string
	GetShellCode() string
}

type FishShell struct{}

func (sh FishShell) GetShellCode(env_value string) string {
	return fmt.Sprintf("set -x PATH %s $PATH", env_value)
}

func (sh FishShell) InitSetup() {
	user_conf_path := sh.GetUserConfigPath()
	user_config_folder := path.Dir(user_conf_path)

	err := os.MkdirAll(user_config_folder, 0755)
	if err != nil {
		panic(err.Error())
	}
	
	// creating user config file in envset folder
	_, err = os.Create(user_conf_path)
	if err != nil {
		panic(err.Error())
	}

	shell_config_file := sh.GetShellConfigPath()
	file, err := os.OpenFile(shell_config_file, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644);
	defer file.Close()
	file.WriteString(fmt.Sprintf("source %s\n", user_conf_path))
}



func (sh FishShell) AddPath()    {
	
}
func (sh FishShell) RemovePath() {}

func (sh FishShell) GetShellConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}

	return path.Join(homeDir, ".config/fish/conf.d/envset.fish")
}

func (sh FishShell) GetUserConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	return path.Join(homeDir, ".config/envset/envset.fish")
}
