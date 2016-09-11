package main

const (
	tokEof = iota + 128
	tokEol
	tokSpace
	tokComment
	tokNumber
	tokString
	tokId

	// commands
	tokData
	tokDef
	tokDim
	tokEnd
	tokFor
	tokGosub
	tokGoto
	tokIf
	tokInput
	tokLet
	tokNext
	tokPrint
	tokRead
	tokRem
	tokRestore
	tokReturn
	tokStop

	// predefined functions
	tokSin
	tokCos
	tokAtn
	tokSqr
	tokExp
	tokLog
	tokAbs
	tokInt
	tokRnd
	tokSgn
	tokVaL
	tokChar
	tokMid
	tokLeft
	tokRight
	tokAsc

	// general
	tokThen
	tokTo
	tokLe
	tokGe
	tokNe
	tokOn
	tokTab
	tokFn
)

type token int

type reservedWord struct {
	token token
	s     string
}

var reservedWords = [...]reservedWord{
	{tokLet, "LET"},
	{tokRead, "READ"},
	{tokData, "DATA"},
	{tokPrint, "PRINT"},
	{tokGoto, "GOTO"},
	{tokIf, "IF"},
	{tokThen, "THEN"},
	{tokFor, "FOR"},
	{tokTo, "TO"},
	{tokNext, "NEXT"},
	{tokDim, "DIM"},
	{tokGosub, "GOSUB"},
	{tokReturn, "RETURN"},
	{tokInput, "INPUT"},
	{tokRem, "REM"},
	{tokRestore, "RESTORE"},
	{tokDef, "DEF"},
	{tokStop, "STOP"},
	{tokEnd, "END"},
	{tokOn, "ON"},
	{tokTab, "TAB"},
	{tokFn, "FN"},
	{tokSin, "SIN"},
	{tokCos, "COS"},
	{tokAtn, "ATN"},
	{tokSqr, "SQR"},
	{tokExp, "EXP"},
	{tokLog, "LOG"},
	{tokAbs, "ABS"},
	{tokInt, "INT"},
	{tokRnd, "RND"},
	{tokSgn, "SGN"},
	{tokVaL, "VAL"},
	{tokChar, "CHAR"},
	{tokMid, "MID"},
	{tokLeft, "LEFT"},
	{tokRight, "RIGHT"},
	{tokAsc, "ASC"},
}
