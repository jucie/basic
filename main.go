package main

import (
	"bufio"
	"os"
)

func compileFromList(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		path = sc.Text()
		compile(path)
	}
}

func compile(path string) {
	if len(path) < 1 {
		return
	}
	if path[0] == '@' {
		compileFromList(path[1:])
		return
	}
	prog := loadProgram(path)
	prog.generate()
}

func main() {
	if len(os.Args) < 2 {
		panic("Missing args")
	}
	for i := 1; i < len(os.Args); i++ {
		path := os.Args[i]
		compile(path)
	}
}
