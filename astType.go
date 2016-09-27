package main

import (
	"fmt"
)

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

type astType interface {
	host
	fmt.Stringer
}

var numType astNumType
var strType astStrType
