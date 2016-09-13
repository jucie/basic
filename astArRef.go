package main

type astArRef struct {
	coord
	id    string
	index []*astExpr
}
