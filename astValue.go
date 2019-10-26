package main

type astValue interface {
	host
	finalType() astType
}
