package main

import (
	"bytes"
)

type generator struct {
	lineAddr map[int]int
	fixups   []int
	text     bytes.Buffer
}

type code struct {
	text   []byte
	fixups []int
}

func (g *generator) consider(h host) {
	/*
		switch v := h.(type) {
		default:
		}
	*/
}

func (g *generator) generate(prog *program) code {
	scan(prog, func(h host) {
		g.consider(h)
	})
	return code{text: g.text.Bytes(), fixups: g.fixups}
}
