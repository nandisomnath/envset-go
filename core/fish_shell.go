package core

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)



type FishShell struct{}

func NewFishShell() FishShell {
	return FishShell{}
}


func (sh FishShell) GetShellCode(env_value string) string {
	return fmt.Sprintf("set -x PATH %s $PATH\n", env_value)
}

func (sh FishShell) InitSetup() {
	userConfPath := sh.GetUserConfigPath()
	userConfigFolder := path.Dir(userConfPath)

	err := os.MkdirAll(userConfigFolder, 0755)
	if err != nil {
		panic(err.Error())
	}
	
	// creating user config file in envset folder
	_, err = os.Create(userConfPath)
	if err != nil {
		panic(err.Error())
	}

	shellConfigFile := sh.GetShellConfigPath()
	file, err := os.OpenFile(shellConfigFile, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644);
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("source %s\n", userConfPath))
}



func (sh FishShell) AddPath(env_value string)    {
	userConfPath := sh.GetUserConfigPath()
	userConfFile, err := os.OpenFile(userConfPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	shell_code := sh.GetShellCode(env_value)
	_, err = userConfFile.WriteString(shell_code)
	if err != nil {
		panic(err.Error())
	}
}

// TODO: handle the unwinding when a error occurs
func (sh FishShell) RemovePath(envValue string) {
	// all variables
	userConfPath := sh.GetUserConfigPath()
	shellCode := sh.GetShellCode(envValue)

	// copy all strings in tmp file
	tmpDir := os.TempDir()
	tmpFile, err := os.CreateTemp(tmpDir, "envset")
	if err != nil {
		panic(err.Error())
	}
	defer tmpFile.Close()
	tmpWriter := bufio.NewWriter(tmpFile)
	userConfFileReaderFile, err := os.Open(userConfPath)
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
	userConfWriterFile, err := os.Create(userConfPath)
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
