package main

import (
	"fmt"
)

type guest interface {
	visit(h host)
}

type host interface {
	receive(g guest)
}

type solver struct {
	p       *program
	dsts    map[int]int
	vars    map[string]int
	types   map[string]int
	funcs   map[string]int
	predefs map[token]int
}

func newSolver(p *program) *solver {
	return &solver{
		p:       p,
		dsts:    make(map[int]int),
		vars:    make(map[string]int),
		types:   make(map[string]int),
		funcs:   make(map[string]int),
		predefs: make(map[token]int),
	}
}

func (s *solver) visit(h host) {
	h.receive(s)
	switch v := h.(type) {
	case *astFnCall:
		s.funcs[v.id]++
	case *astPredef:
		s.predefs[v.function]++
	case astType:
		t := fmt.Stringer(v)
		s.types[t.String()]++
	case astVarRef:
		s.vars[v.id]++
	case cmdDef:
		s.funcs[v.id]++
	case cmdGoto:
		s.dsts[v.dst]++
	case cmdGosub:
		s.dsts[v.dst]++
	case cmdOn:
		for _, dst := range v.dsts {
			s.dsts[dst]++
		}
	}
}

func (s *solver) showStats() {
	if len(s.vars) > 0 {
		println("\nVars")
		for key, val := range s.vars {
			println("\t", key, val)
		}
	}

	if len(s.dsts) > 0 {
		println("\nDsts")
		for key, val := range s.dsts {
			println("\t", key, val)
		}
	}

	if len(s.types) > 0 {
		println("\nTypes")
		for key, val := range s.types {
			println("\t", key, val)
		}
	}

	if len(s.funcs) > 0 {
		println("\nFunctions")
		for key, val := range s.funcs {
			println("\t", key, val)
		}
	}

	if len(s.predefs) > 0 {
		println("\nPredefs")
		for key, val := range s.predefs {
			println("\t", key, val)
		}
	}
}
