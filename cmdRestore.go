package main

type cmdRestore struct {
}

func (p *parser) parseRestore() *cmdRestore {
	return &cmdRestore{}
}
