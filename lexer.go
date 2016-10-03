package main

import (
	"bufio"
	"bytes"
	"strings"
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
	rd    *bufio.Reader
	pos   coord
	begin coord
	buf   bytes.Buffer
	lexeme
	unreadBuf bytes.Buffer
	c         chan lexeme
}

func newLexer(rd *bufio.Reader) *lexer {
	lex := &lexer{rd: rd}
	lex.c = make(chan lexeme)
	go func() {
		for {
			lex.next0()
		}
	}()
	lex.next()
	return lex
}

func (lex *lexer) get() *lexeme {
	l := lex.peek()
	lex.next()
	return l
}

func (lex *lexer) peek() *lexeme {
	return &lex.lexeme
}

func (lex *lexer) next() {
	lex.lexeme = <-lex.c
}

func (lex *lexer) next0() {
	var tok token
	for {
		tok = lex.walk()
		if tok == tokRem || tok == tokData {
			lex.consumeLine()
		}
		lex.pos.col += lex.buf.Len()
		if tok != tokSpace {
			break
		}
	}
	lex.c <- lexeme{token: tok, pos: lex.pos, s: string(lex.buf.Bytes())}
}

func (lex *lexer) nextLine() {
	lex.consumeLine()
	if lex.peek().token == tokEol {
		lex.walk()
	}
}

func (lex *lexer) consumeLine() {
	lex.buf.Reset()

	for {

		b, err := lex.readByte()
		if err != nil {
			break
		}

		if b == '\n' || b == '\r' {
			lex.unreadByte(b)
			break
		}
		lex.buf.WriteByte(b)
	}
}

func (lex *lexer) handleSpace(b byte) token {
	for {
		lex.buf.WriteByte(b)
		b, err := lex.readByte()
		if err != nil {
			break
		}
		if b == '\n' || !unicode.IsSpace(rune(b)) {
			lex.unreadByte(b)
			break
		}
	}
	return tokSpace
}

func (lex *lexer) handleString(b byte) token {
	lex.buf.WriteByte(b)
	for {
		b, err := lex.readByte()
		if err != nil {
			break
		}
		if b == '\n' {
			lex.unreadByte(b)
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
	lex.buf.WriteByte(byte(unicode.ToUpper(rune(b))))

	var err error
	b, err = lex.readByte()
	if err != nil {
		return tokId
	}

	// for an id, if the second character is a digit, then we are done.
	if unicode.IsNumber(rune(b)) {
		lex.buf.WriteByte(byte(unicode.ToUpper(rune(b))))
		return tokId // we have a single character id
	}

	// for an id, the second character must be a letter or a digit
	if !unicode.IsLetter(rune(b)) {
		lex.unreadByte(b)
		return tokId // we have a single character id
	}

	var lastByte byte
	// read full word
	for {
		lex.buf.WriteByte(byte(unicode.ToUpper(rune(b))))
		b, err = lex.readByte()
		if err != nil {
			break
		}
		if !unicode.IsLetter(rune(b)) {
			lastByte = b
			break
		}
	}

	// trying to find keywords in the string.
	s := string(lex.buf.Bytes())
	min := len(s)
	var rwmin reservedWord
	for _, rw := range reservedWords {
		pos := strings.Index(s, rw.s)
		if pos < 0 {
			continue
		} else if pos == 0 {
			min = 0
			rwmin = rw
			break
		} else {
			if pos < min {
				min = pos // the earlier find is what we want.
				rwmin = rw
			}
		}
	}
	var tok token = tokId
	if min == 0 { // the string begins with a keyword
		lex.buf.Reset()
		lex.buf.WriteString(rwmin.s)
		lex.unreadString(s[len(rwmin.s):])
		tok = rwmin.token
	} else if min > 0 && min < len(s) { // we have an id before the keyword
		lex.buf.Reset()
		lex.buf.WriteString(s[:min])
		lex.unreadString(s[min:])
	}
	lex.unreadByte(lastByte) // the non letter character after the string
	return tok
}

func (lex *lexer) handleDigraph(b byte) token {
	lex.buf.WriteByte(b)

	second, err := lex.readByte()
	if err != nil {
		return token(b)
	}

	lex.buf.WriteByte(second)
	switch string(lex.buf.Bytes()) {
	case "<=":
		return tokLe
	case ">=":
		return tokGe
	case "<>":
		return tokNe
	}

	lex.unreadByte(second) // unread second
	lex.buf.Reset()
	lex.buf.WriteByte(b)
	return token(b)
}

func (lex *lexer) handleNumber(b byte) token {
	var err error
	hasPoint := false
	hasE := false
	hasExpSignal := false
Loop:
	for {
		switch b {
		case '.':
			if hasPoint {
				lex.unreadByte(b)
				break Loop
			}
			hasPoint = true
			lex.buf.WriteByte(b)
		case 'E':
			if hasE {
				lex.unreadByte(b)
				break Loop
			}
			hasE = true
			lex.buf.WriteByte(b)
		case '+':
			fallthrough
		case '-':
			if !hasE {
				lex.unreadByte(b)
				break Loop
			}
			if hasExpSignal {
				lex.unreadByte(b)
				break Loop
			}
			hasExpSignal = true
			lex.buf.WriteByte(b)
		default:
			if !unicode.IsDigit(rune(b)) {
				lex.unreadByte(b)
				break Loop
			}
			lex.buf.WriteByte(b)
		}
		b, err = lex.readByte()
		if err != nil {
			break Loop
		}
	}
	return tokNumber
}

func canBeNumber(b byte) bool {
	return b == '.' || unicode.IsDigit(rune(b))
}

func (lex *lexer) walk() token {
	lex.buf.Reset()

	b, err := lex.readByte()
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

	if canBeNumber(b) {
		return lex.handleNumber(b)
	}

	if unicode.IsLetter(rune(b)) {
		return lex.handleId(b)
	}

	if b == '"' {
		return lex.handleString(b)
	}

	if b == '>' || b == '<' {
		return lex.handleDigraph(b)
	}

	lex.buf.WriteByte(b)
	return token(b)
}

func (lex *lexer) readByte() (byte, error) {
	if lex.unreadBuf.Len() > 0 {
		return lex.unreadBuf.ReadByte()
	}
	return lex.rd.ReadByte()
}

func (lex *lexer) unreadByte(b byte) {
	lex.unreadBuf.WriteByte(b)
}

func (lex *lexer) unreadString(s string) {
	lex.unreadBuf.WriteString(s)
}
