package main

import (
	"bufio"
	"os"
	"strings"
)

type cmd interface {
}

type progLine struct {
	id   int
	cmds []cmd
}

type program struct {
	srcPath string
	dstPath string
	lines   map[int]*progLine
}

func newProgram() *program {
	return &program{lines: make(map[int]*progLine, 0)}
}

func loadProgram(path string) *program {
	prog := newProgram()

	prog.srcPath = path
	pos := strings.LastIndexByte(path, '.')
	if pos < 0 {
		pos = len(path)
	}
	prog.dstPath = path[:pos] + ".exe"

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	newParser(rd).parseProgram(prog)
	return prog
}

func (prog *program) generate() {
	// TODO
}
