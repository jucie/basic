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
	previous  lexeme
	ids       map[string]bool
	unreadBuf bytes.Buffer
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
	lex := &lexer{rd: rd, ids: make(map[string]bool)}
	lex.pos.row = 1
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
		if l.token == tokRem || l.token == tokData {
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

	var firstNonLetter byte
	// read the longest letter string
	for {
		lex.buf.WriteByte(byte(unicode.ToUpper(rune(b))))
		b, err = lex.readByte()
		if err != nil {
			break
		}
		if !unicode.IsLetter(rune(b)) {
			firstNonLetter = b
			break
		}
	}

	// search for keyword as prefix
	s := string(lex.buf.Bytes())
	for keyword, tok := range keywordMap {
		if strings.HasPrefix(s, keyword) {
			lex.unreadString(s[len(keyword):])
			lex.unreadByte(firstNonLetter)
			lex.buf.Reset()
			lex.buf.WriteString(keyword)
			return tok
		}
	}

	// search for keyword as suffix
	var unread string
	for {
		max := 0
		index := len(s)
		for keyword, _ := range keywordMap {
			if strings.HasSuffix(s, keyword) {
				if max < len(keyword) {
					max = len(keyword)
				}
			}
		}
		if max > 0 {
			index -= max
			keyword := s[index:]
			unread = keyword + unread
			s = s[:index]
		} else {
			break
		}
	}

	if len(unread) > 0 {
		lex.unreadString(unread)
		lex.buf.Reset()
		lex.buf.WriteString(s)
	}
	lex.unreadByte(firstNonLetter)
	return tokId
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

func (lex *lexer) addId(id string) {
	lex.ids[id] = true
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
