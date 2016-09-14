package main

type astTerm interface {
	astPart
}

type astTermImpl struct {
	astPartImpl
	coord
}
