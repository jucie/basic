package main

const (
	tokEof = iota
	tokSpace
	tokComment
	tokInt
	tokFloat
	tokString
	tokId
)

type token int
