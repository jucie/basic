package main

import (
	"io"
)

type lexer struct {
	rd io.Reader
}

func newLexer(rd io.Reader) *lexer {
	return &lexer{rd: rd}
}
