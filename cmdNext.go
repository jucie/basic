package main

type cmdNext struct {
}

func (p *parser) parseNext() *cmdNext {
	return &cmdNext{}
}
