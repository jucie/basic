package main

import (
	"bufio"
)

type generatorForC struct {
}

func newGeneratorForC() *generatorForC {
	return &generatorForC{}
}

func (gen *generatorForC) generate(wr *bufio.Writer, prog *program) {
	prog.generateC(wr)
}
