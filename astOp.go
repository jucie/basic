package main

type astOp interface {
	astPart
	Lhs() astPart
	Rhs() astPart
}

type astOpImpl struct {
	astPartImpl
	lsh astPart
	rhs astPart
}
