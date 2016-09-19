package main

type astMulOpTail struct {
	oper token
	val  *astExpOp
}

type astMulOp struct {
	head *astExpOp
	tail []astMulOpTail
}

func (p *parser) parseMulOp() *astMulOp {
	println(">parseMulOp")
	defer println("<parseMulOp")

	head := p.parseExpOp()
	result := &astMulOp{head: head}

	for {
		oper := p.lex.peek().token
		if !isMulOp(oper) {
			break
		}
		p.lex.next()

		val := p.parseExpOp()
		if val == nil {
			return nil
		}
		result.tail = append(result.tail, astMulOpTail{oper: oper, val: val})
	}
	return result
}

func isMulOp(b token) bool {
	return b == '*' || b == '/'
}
