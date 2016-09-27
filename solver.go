package main

type guest interface {
	visit(h host)
}

type host interface {
	receive(g guest)
}

type solver struct {
	p *program
}

func newSolver(p *program) *solver {
	return &solver{p: p}
}

func (s *solver) visit(h host) {
	if h == nil {
		return
	}
	h.receive(s)
}
