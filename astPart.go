package main

type astPart interface {
	Signal() bool
}

type astPartImpl struct  {
	signal bool
}
