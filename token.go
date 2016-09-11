package main

const (
	tokEof = iota
	tokEol
	tokSpace
	tokComment
	tokInt
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
}
