package main

const (
	tokNil = iota
	tokSpace
	tokComment
	tokInt
	tokFloat
	tokString
	tokId
)

type token int
