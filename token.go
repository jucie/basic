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

var reservedWords = [...]reservedWord{
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
	{tokGosub, "GOSUB"},
	{tokGoto, "GOTO"},
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
	{tokTab, "TAB"},
	{tokTan, "TAN"},
	{tokThen, "THEN"},
	{tokTo, "TO"},
	{tokVal, "VAL"},
}
