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
	previous lexeme
}

var keywordMap = make(map[string]token)

func init() {
	for i, _ := range reservedWords {
		word := &reservedWords[i]
		_, ok := keywordMap[word.s]
		if ok {
			panic("Repeated keyword: " + word.s)
		}
		keywordMap[word.s] = word.token
	}
}

func newLexer(rd *bufio.Reader) *lexer {
	lex := &lexer{rd: rd}
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
	lex.previous = lex.lexeme
	l := &lex.lexeme
	for {
		l.pos = lex.pos
		l.token = lex.walk()
		if l.token == tokRem {
			lex.consumeLine()
		}
		lex.pos.col += lex.buf.Len()
		if l.token != tokSpace {
			break
		}
	}
	l.s = string(lex.buf.Bytes())
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

		b, err := lex.rd.ReadByte()
		if err != nil {
			break
		}

		if b == '\n' || b == '\r' {
			lex.rd.UnreadByte()
			break
		}
		lex.buf.WriteByte(b)
	}
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
	lex.buf.WriteByte(byte(unicode.ToUpper(rune(b))))

	second, err := lex.rd.ReadByte()
	if err != nil {
		return tokId
	}

	// for an id, the second character must be a letter, a digit or a dollar sign.
	if !unicode.IsLetter(rune(second)) && !unicode.IsNumber(rune(second)) && second != '$' {
		lex.rd.UnreadByte()
		return tokId // we have a single character id
	}

	lex.buf.WriteByte(byte(unicode.ToUpper(rune(second))))
	s := string(lex.buf.Bytes())

	tok, ok := keywordMap[s]
	if ok { // if it is a two letter keyword like IF or ON
		return tok
	}

	// let's see if it can be the beginning of a keyword
	found := false
	for key, _ := range keywordMap {
		if strings.HasPrefix(key, s) {
			found = true
			break
		}
	}
	if !found { // if it can't possibly start a keyword
		return tokId
	}

	// ok, it MUST be a keyword
	for {
		var err error
		b, err = lex.rd.ReadByte()
		if err != nil {
			break
		}
		if !unicode.IsLetter(rune(b)) && b != '$' {
			lex.rd.UnreadByte()
			break
		}
		lex.buf.WriteByte(byte(unicode.ToUpper(rune(b))))
		s := string(lex.buf.Bytes())
		tok, ok := keywordMap[s]
		if ok {
			return tok
		}
	}
	return tokId
}

func (lex *lexer) handleDigraph(b byte) token {
	lex.buf.WriteByte(b)

	second, err := lex.rd.ReadByte()
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

	lex.rd.UnreadByte() // unread second
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
				lex.rd.UnreadByte()
				break Loop
			}
			hasPoint = true
			lex.buf.WriteByte(b)
		case 'E':
			if hasE {
				lex.rd.UnreadByte()
				break Loop
			}
			hasE = true
			lex.buf.WriteByte(b)
		case '+':
			fallthrough
		case '-':
			if !hasE {
				lex.rd.UnreadByte()
				break Loop
			}
			if hasExpSignal {
				lex.rd.UnreadByte()
				break Loop
			}
			hasExpSignal = true
			lex.buf.WriteByte(b)
		default:
			if !unicode.IsDigit(rune(b)) {
				lex.rd.UnreadByte()
				break Loop
			}
			lex.buf.WriteByte(b)
		}
		b, err = lex.rd.ReadByte()
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
