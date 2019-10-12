package main

import "bufio"

type astValue interface {
	host
	generateC(wr *bufio.Writer)
	finalType() astType
}
