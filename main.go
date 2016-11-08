package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	prog.resolve()
	prog.generate()
	fmt.Printf("Parsed program %s (%d lines, %d blocks, %d orphans)\n", prog.srcPath, len(prog.lines), len(prog.blocks), prog.orphans)
}

func main() {
	if len(os.Args) < 2 {
		panic("Missing args")
	}
	for i := 1; i < len(os.Args); i++ {
		path := os.Args[i]
		compile(path)
	}
	var typenames []string
	for key, _ := range mt {
		typenames = append(typenames, key)
	}
	sort.Strings(typenames)
	for _, typename := range typenames {
		fmt.Printf("%6d\t%s\n", mt[typename], typename)
	}
}
