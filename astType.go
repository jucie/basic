package main

import (
	"fmt"
)

type astType interface {
	host
	fmt.Stringer
}

type astNumType struct {
}

func (a astNumType) String() string {
	return "num"
}

func (a astNumType) receive(g guest) {
}

type astStrType struct {
}

func (a astStrType) String() string {
	return "str"
}

func (a astStrType) receive(g guest) {
}

type astVoidType struct {
}

func (a astVoidType) String() string {
	return "void"
}

func (a astVoidType) receive(g guest) {
}

var numType astNumType
var strType astStrType
var voidType astVoidType
