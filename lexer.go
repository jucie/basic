package main

import (
	"bufio"
	"bytes"
	"unicode"
)

type coord struct {
	row int
	col int
}

type lexeme struct {
	token token
	pos   coord
	s     string
}

type lexer struct {
	rd     *bufio.Reader
	pos    coord
	begin  coord
	buf    bytes.Buffer
	lexeme lexeme
}

var keywordMap = make(map[string]token)

func init() {
	for i, _ := range reservedWords {
		word := &reservedWords[i]
		keywordMap[word.s] = word.token
	}
}

func newLexer(rd *bufio.Reader) *lexer {
	lex := &lexer{rd: rd}
	lex.next()
	return lex
}

func (lex *lexer) peek() *lexeme {
	return &lex.lexeme
}

func (lex *lexer) next() {
	l := &lex.lexeme
	for {
		l.pos = lex.pos
		l.token = lex.walk()
		lex.pos.col += lex.buf.Len()
		if l.token != tokSpace {
			break
		}
	}
	l.s = string(lex.buf.Bytes())
}

func (lex *lexer) handleSpace(b byte) token {
	for {
		lex.buf.WriteByte(b)
		b, err := lex.rd.ReadByte()
		if err != nil {
			break
		}
		if b == '\n' || !unicode.IsSpace(rune(b)) {
			lex.rd.UnreadByte()
			break
		}
	}
	return tokSpace
}

func (lex *lexer) handleNumber(b byte) token {
	for {
		lex.buf.WriteByte(b)
		var err error
		b, err = lex.rd.ReadByte()
		if err != nil {
			break
		}
		if !unicode.IsDigit(rune(b)) {
			lex.rd.UnreadByte()
			break
		}
	}
	return tokInt
}

func (lex *lexer) handleString(b byte) token {
	lex.buf.WriteByte(b)
	for {
		b, err := lex.rd.ReadByte()
		if err != nil {
			break
		}
		if b == '\n' {
			lex.rd.UnreadByte()
			break
		}
		lex.buf.WriteByte(b)
		if b == '"' {
			break
		}
	}
	return tokString
}

func (lex *lexer) handleId(b byte) token {
	for {
		lex.buf.WriteByte(b)
		s := string(lex.buf.Bytes())
		tok, ok := keywordMap[s]
		if ok {
			return tok
		}

		var err error
		b, err = lex.rd.ReadByte()
		if err != nil {
			break
		}
		if !unicode.IsLetter(rune(b)) && !unicode.IsNumber(rune(b)) {
			break
		}
	}
	return tokId
}

func (lex *lexer) walk() token {
	lex.buf.Reset()
	b, err := lex.rd.ReadByte()
	if err != nil {
		return tokEof
	}

	if b == '\n' {
		lex.pos.row++
		lex.pos.col = 0
		return tokEol
	}

	if unicode.IsSpace(rune(b)) {
		return lex.handleSpace(b)
	}

	if unicode.IsDigit(rune(b)) || b == '-' || b == '+' {
		return lex.handleNumber(b)
	}

	if unicode.IsLetter(rune(b)) {
		return lex.handleId(b)
	}

	if b == '"' {
		return lex.handleString(b)
	}

	lex.buf.WriteByte(b)
	return token(b)
}
