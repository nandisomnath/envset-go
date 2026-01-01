package main

import "os"


func main() {
	os.MkdirAll("./hello", 0755)
}