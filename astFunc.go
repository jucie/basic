package main

type astFunc struct {
	coord
	id   rune
	arg  rune
	expr *astExpr
}
