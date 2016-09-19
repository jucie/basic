package main

type astOp struct {
	lhs  *astPart
	rhs  *astPart
	oper token
}
