package main

import (
	"bufio"
	"fmt"
	"os"
)

type program struct {
	source []string
}

func NewProgram() *program {
	return new(program)
}

func (prog *program) load(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		prog.source = append(prog.source, line)
	}
}

func (prog *program) generate() {
	for _, line := range prog.source {
		fmt.Println(line)
	}
}
