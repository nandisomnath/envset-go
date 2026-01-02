package core


func BashShellHandler(envValue string, removePath bool)  {
	// TODO: implement it
}

func ZshShellHandler(envValue string, removePath bool)  {
	// TODO: implement it
}

func FishShellHandler(envValue string, removePath bool)  {
	shell := NewFishShell()

	if removePath {
		shell.RemovePath(envValue)
	} else {
		shell.AddPath(envValue)
	}
}