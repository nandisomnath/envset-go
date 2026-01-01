package core

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
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
	return fmt.Sprintf("set -x PATH %s $PATH\n", env_value)
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



func (sh FishShell) AddPath(env_value string)    {
	user_conf_path := sh.GetUserConfigPath()
	user_conf_file, err := os.OpenFile(user_conf_path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	shell_code := sh.GetShellCode(env_value)
	_, err = user_conf_file.WriteString(shell_code)
	if err != nil {
		panic(err.Error())
	}
}

func (sh FishShell) RemovePath(env_value string) {
	// all variables
	user_conf_path := sh.GetUserConfigPath()
	shellCode := sh.GetShellCode(env_value)

	// copy all strings in tmp file
	tmpDir := os.TempDir()
	tmpFile, err := os.CreateTemp(tmpDir, "envset")
	if err != nil {
		panic(err.Error())
	}
	defer tmpFile.Close()
	tmpWriter := bufio.NewWriter(tmpFile)
	userConfFileReaderFile, err := os.Open(user_conf_path)
	if err != nil {
		panic(err.Error())
	}
	defer userConfFileReaderFile.Close()
	userConfFileScanner := bufio.NewScanner(userConfFileReaderFile)

	for userConfFileScanner.Scan() {
		line := userConfFileScanner.Text()
		if strings.TrimSpace(shellCode) != line {
			_, err = tmpWriter.WriteString(fmt.Sprintf("%s\n", line))
			if err != nil {
				panic(err.Error())
			}
		}
	}

	// read from temp file and write it in userConf file
	tmpFileReader := bufio.NewScanner(tmpFile)
	userConfWriterFile, err := os.Create(user_conf_path)
	if err != nil {
		panic(err.Error())
	}
	defer userConfWriterFile.Close()
	userConfFileWriter := bufio.NewWriter(userConfWriterFile)
	for tmpFileReader.Scan() {
		line := tmpFileReader.Text()
		userConfFileWriter.WriteString(fmt.Sprintf("%s\n", line))
	}

}

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
