package main

type astNumType struct {
}

func (a astNumType) receive(g guest) {
}

type astStrType struct {
}

func (a astStrType) receive(g guest) {
}

type astType interface {
	host
}

var numType astNumType
var strType astStrType
