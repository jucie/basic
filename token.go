package main

import (
	"sort"
)

const (
	tokEOF = iota + 128
	tokEOL
	tokSpace
	tokComment
	tokNumber
	tokString
	tokID

	// commands
	tokData
	tokDef
	tokDim
	tokEnd
	tokFor
	tokGo
	tokIf
	tokInput
	tokLet
	tokNext
	tokOn
	tokPrint
	tokRead
	tokRem
	tokRestore
	tokReturn
	tokRun
	tokStop

	// predefined functions
	tokAbs
	tokAsc
	tokAtn
	tokChr
	tokCos
	tokExp
	tokInt
	tokLeft
	tokLen
	tokLog
	tokMid
	tokRight
	tokRnd
	tokSgn
	tokSin
	tokSqr
	tokStr
	tokTab
	tokTan
	tokVal

	// general
	tokThen
	tokTo
	tokSub
	tokStep
	tokLe
	tokGe
	tokNe
	tokFn
	tokAnd
	tokOr
)

type token int

func (t token) receive(g guest) {
}

type reservedWord struct {
	token token
	s     string
}

var reservedWordList = [...]reservedWord{
	{tokAbs, "ABS"},
	{tokAnd, "AND"},
	{tokAsc, "ASC"},
	{tokAtn, "ATN"},
	{tokChr, "CHR"},
	{tokCos, "COS"},
	{tokData, "DATA"},
	{tokDef, "DEF"},
	{tokDim, "DIM"},
	{tokEnd, "END"},
	{tokExp, "EXP"},
	{tokFn, "FN"},
	{tokFor, "FOR"},
	{tokGo, "GO"},
	{tokIf, "IF"},
	{tokInput, "INPUT"},
	{tokInt, "INT"},
	{tokLeft, "LEFT"},
	{tokLen, "LEN"},
	{tokLet, "LET"},
	{tokLog, "LOG"},
	{tokMid, "MID"},
	{tokNext, "NEXT"},
	{tokOn, "ON"},
	{tokOr, "OR"},
	{tokPrint, "PRINT"},
	{tokRead, "READ"},
	{tokRem, "REM"},
	{tokRestore, "RESTORE"},
	{tokReturn, "RETURN"},
	{tokRight, "RIGHT"},
	{tokRnd, "RND"},
	{tokRun, "RUN"},
	{tokSgn, "SGN"},
	{tokSin, "SIN"},
	{tokSqr, "SQR"},
	{tokStep, "STEP"},
	{tokStop, "STOP"},
	{tokStr, "STR"},
	{tokSub, "SUB"},
	{tokTab, "TAB"},
	{tokTan, "TAN"},
	{tokThen, "THEN"},
	{tokTo, "TO"},
	{tokVal, "VAL"},
}

type rws []reservedWord

func (l rws) Len() int { return len(l) }
func (l rws) Less(i, j int) bool {
	lis := l[i].s
	ljs := l[j].s
	if len(lis) > len(ljs) {
		return true
	}
	if len(lis) < len(ljs) {
		return false
	}
	return lis > ljs
}
func (l rws) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

var reservedWords rws

func init() {
	reservedWords = reservedWordList[:]
	sort.Sort(reservedWords)
}
