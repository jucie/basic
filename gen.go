package main

import (
	"bytes"
)

type generator struct {
	lineAddr map[int]int
	fixups   []int
	text     bytes.Buffer
	data     []string
}

type code struct {
	text   []byte
	data   []string
	lit    []string
	fixups []int
}

func (g *generator) emitByte(b byte) {
	g.text.WriteByte(b)
}

func (g *generator) emitBytes(val ...byte) {
	for _, b := range val {
		g.emitByte(b)
	}
}

func (g *generator) consider(h host) {
	switch v := h.(type) {
	case *cmdData:
		g.data = append(g.data, v.values...)
	case *cmdEnd:
		g.emitBytes(0x5F, 0x5D, 0xC3) // pop rdi; pop rbp; ret
	case *cmdStop:
		g.emitBytes(0x5F, 0x5D, 0xC3) // pop rdi; pop rbp; ret
	default:
	}
}

func (g *generator) generate(prog *program) code {
	scan(prog, func(h host) {
		g.consider(h)
	})
	return code{text: g.text.Bytes(), data: g.data, fixups: g.fixups}
}
