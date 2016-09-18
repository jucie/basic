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
	tokStep
	tokLe
	tokGe
	tokNe
	tokTab
	tokFn
	tokAnd
	tokOr
)

type token int

type reservedWord struct {
	token token
	s     string
}

var reservedWords = [...]reservedWord{
	{tokAbs, "ABS"},
	{tokAnd, "AND"},
	{tokAsc, "ASC"},
	{tokAtn, "ATN"},
	{tokChar, "CHAR"},
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
	{tokSgn, "SGN"},
	{tokAbs, "ABS"},
	{tokAsc, "ASC"},
	{tokAtn, "ATN"},
	{tokChar, "CHAR"},
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
	{tokLet, "LET"},
	{tokLog, "LOG"},
	{tokMid, "MID"},
	{tokNext, "NEXT"},
	{tokOn, "ON"},
	{tokPrint, "PRINT"},
	{tokRead, "READ"},
	{tokRem, "REM"},
	{tokRestore, "RESTORE"},
	{tokReturn, "RETURN"},
	{tokRight, "RIGHT"},
	{tokRnd, "RND"},
	{tokSgn, "SGN"},
	{tokSin, "SIN"},
	{tokSqr, "SQR"},
	{tokStop, "STOP"},
	{tokTab, "TAB"},
	{tokThen, "THEN"},
	{tokTo, "TO"},
	{tokVaL, "VAL"},
	{tokSin, "SIN"},
	{tokSqr, "SQR"},
	{tokStep, "STEP"},
	{tokStop, "STOP"},
	{tokTab, "TAB"},
	{tokThen, "THEN"},
	{tokTo, "TO"},
	{tokVaL, "VAL"},
}
