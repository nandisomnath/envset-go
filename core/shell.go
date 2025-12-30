package core

type Shell interface {
	RemovePath()
	AddPath()
	InitSetup()
	GetShellConfigPath() string
	GetUserConfigPath() string
	GetShellCode() string
}

type FishShell struct{}

func (sh FishShell) GetShellCode() string {
	return ""
}

func (sh FishShell) InitSetup()                 {}
func (sh FishShell) AddPath()                   {}
func (sh FishShell) RemovePath()                {}
func (sh FishShell) GetShellConfigPath() string { return "" }
func (sh FishShell) GetUserConfigPath() string  { return "" }
