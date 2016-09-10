package main

import (
	"os"
)

func loadSourceFile(path string) {

}

func generateExecutable() {

}

func main() {
	if len(os.Args) < 2 {
		panic("Missing args")
	}
	for i := 1; i < len(os.Args); i++ {
		path := os.Args[i]
		loadSourceFile(path)
	}
	generateExecutable()
}
