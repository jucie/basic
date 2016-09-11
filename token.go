package main

const (
	tokEof = iota + 128
	tokEol
	tokSpace
	tokComment
	tokNumber
	tokString
	tokId

	tokLet
	tokRead
	tokData
	tokPrint
	tokGoto
	tokIf
	tokThen
	tokFor
	tokTo
	tokNext
	tokDim
	tokGosub
	tokReturn
	tokInput
	tokRem
	tokRestore
	tokDef
	tokStop
	tokEnd
	tokLe
	tokGe
	tokNe
	tokOn
	tokTab
	tokFn
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
