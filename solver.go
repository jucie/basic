package main

import (
	"fmt"
	"os"
)

type variable struct {
	def        *astVarRef
	dims       int
	ref        []*astVarRef
	isForIndex bool
}

type function struct {
	id  string
	def *cmdFnDef
	ref []*astFnCall
}

type solver struct {
	p       *program
	dsts    []*targetLine
	types   map[string]int
	funcs   map[string]*function
	predefs map[token]int
	vars    map[string]*variable
}

func newSolver(p *program) *solver {
	return &solver{
		p:       p,
		types:   make(map[string]int),
		funcs:   make(map[string]*function),
		predefs: make(map[token]int),
		vars:    make(map[string]*variable),
	}
}

var mt = make(map[string]int)

func (s *solver) consider(h host) {
	mt[fmt.Sprintf("%T", h)]++
	switch v := h.(type) {
	case *astPredef:
		s.predefs[v.function]++
	case astType:
		t := fmt.Stringer(v)
		s.types[t.String()]++
	case *astVarRef:
		vv, ok := s.vars[v.id]
		if !ok {
			vv = &variable{}
			s.vars[v.id] = vv
		}
		vv.ref = append(vv.ref, v)
		vv.dims = len(v.index)
	case *cmdDim:
		for _, def := range v.vars {
			vv, ok := s.vars[def.nameForC()]
			if !ok {
				vv = &variable{}
				s.vars[def.nameForC()] = vv
			}
			if vv.def != nil {
				fmt.Fprintf(os.Stderr, "Multiple definition for variable %s.", def.id)
			}
			vv.def = def
			vv.dims = len(def.index)
		}
	case *cmdFor:
		vv, ok := s.vars[v.index.id]
		if !ok {
			vv = &variable{}
			s.vars[v.index.id] = vv
		}
		vv.isForIndex = true
	case *cmdFnDef:
		vv, ok := s.funcs[v.id]
		if !ok {
			vv = &function{id: v.id}
			s.funcs[v.id] = vv
		}
		if vv.def != nil {
			fmt.Fprintf(os.Stderr, "Multiple definition for function %s.", v.id)
		}
		vv.def = v
	case *astFnCall:
		vv, ok := s.funcs[v.id]
		if !ok {
			vv = &function{id: v.id}
			s.funcs[v.id] = vv
		}
		vv.ref = append(vv.ref, v)
	case *cmdGo:
		s.dsts = append(s.dsts, &v.dst)
	case *cmdOn:
		for i := 0; i < len(v.dsts); i++ {
			s.dsts = append(s.dsts, &v.dsts[i])
		}
	}
}

func (s *solver) linkLines(lines progLines) {
	m := make(map[int]*progLine)
	for _, l := range lines {
		m[l.id] = l
	}
	for i := range s.dsts {
		dst := s.dsts[i]
		l, ok := m[dst.nbr]
		if !ok {
			l = lines.find(dst.nbr)
		}
		if l == nil {
			fmt.Fprintf(os.Stderr, "Target line not found: %d\n", dst.nbr)
		} else {
			dst.adr = l
			l.isDst = true
		}
	}
}
