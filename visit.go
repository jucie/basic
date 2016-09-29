package main

type guest interface {
	visit(h host)
}

type host interface {
	receive(g guest)
}

type visitor struct {
	consider func(h host)
}

func (v *visitor) visit(h host) {
	h.receive(v)
	v.consider(h)
}

func scan(h host, consider func(h host)) {
	v := visitor{consider: consider}
	v.visit(h)
}
