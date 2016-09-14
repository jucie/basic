package main

type astVarRef struct {
	astTermImpl
	id    string
	index []*astExpr
}
