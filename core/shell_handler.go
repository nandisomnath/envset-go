package core


func BashShellHandler(envValue string, removePath bool)  {
	
}

func ZshShellHandler(envValue string, removePath bool)  {
	
}

func FishShellHandler(envValue string, removePath bool)  {
	shell := NewFishShell()

	if removePath {
		shell.RemovePath(envValue)
	} else {
		shell.AddPath(envValue)
	}
}