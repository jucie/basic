package main

type astFnCall struct {
	coord
	fn  rune
	arg *astExpr
}
