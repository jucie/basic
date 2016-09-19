package main

type astExpOp struct {
	head *astPart
	tail []*astPart
}

func (p *parser) parseExpOp() *astExpOp {
	head := p.parsePart()
	result := &astExpOp{head: head}

	for {
		oper := p.lex.peek().token
		if oper != '^' {
			break
		}
		p.lex.next()

		val := p.parsePart()
		if val == nil {
			return nil
		}
		result.tail = append(result.tail, val)
	}
	return result
}
