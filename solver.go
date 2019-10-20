package main

import (
	"fmt"
	"os"
	"sort"
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
			vv, ok := s.vars[def.unambiguousName()]
			if !ok {
				vv = &variable{}
				s.vars[def.unambiguousName()] = vv
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

func (s *solver) linkLines(lines progLines) progLines {
	m := make(map[int]*progLine)
	for _, l := range lines {
		m[l.id] = l
	}
	for i := range s.dsts {
		dst := s.dsts[i]
		l, ok := m[dst.nbr]
		if !ok { // target line doesn't exist
			l = &progLine{id: dst.nbr}
			m[dst.nbr] = l
			lines = append(lines, l)
		}
		dst.adr = l
		l.isDst = true
	}
	sort.Slice(lines, func(i, j int) bool { return lines[i].id < lines[j].id })
	return lines
}

func (s *solver) linkForNext(p *program) {
	m := make(map[string]struct{})
	stack := make([]*cmdFor, 64) // max levels deep
	sp := -1
	scan(p, func(h host) {
		switch v := h.(type) {
		case *cmdFor:
			sp++
			if sp >= len(stack) {
				panic("Too many nested FOR loops.")
			}
			stack[sp] = v
			m[v.index.unambiguousName()] = struct{}{}
		case *cmdNext:
			v.createLabel()
			if sp < 0 {
				panic("NEXT without FOR")
			}
			if len(v.vars) == 0 { // NEXT without a var only matches the most recent FOR
				f := stack[sp]
				f.next = v
				v.vars = append(v.vars, f.index)
				sp--
			}
			// NEXT with a var
			for sp >= 0 {
				f := stack[sp]
				if !v.vars[0].equals(f.index) { // FOR doesn't match NEXT
					break
				}
				f.next = v
				v.vars = append(v.vars, f.index)
				sp--
			}
		}
	})
	v := make([]string, 0, len(m))
	for k := range m {
		v = append(v, k)
	}
	sort.Strings(v)
	p.loopVars = v
}
