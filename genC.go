package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

type generatorForC struct {
	wr *bufio.Writer
}

func newGeneratorForC(wr *bufio.Writer) *generatorForC {
	return &generatorForC{wr: wr}
}

func (g *generatorForC) generate(prog *program) {
	g.genProgram(prog)
	/*
		scan(p, func(h host) {
			switch v := h.(type) {
			case *cmdFor:
			case *cmdNext:
			}
		}
	*/
}

func (g *generatorForC) genProgram(p *program) {
	fmt.Fprintf(g.wr, "#include \"basiclib.h\"\n\n")
	g.genPrologue(p)
	fmt.Fprintf(g.wr, "int main(){\n")
	for _, s := range p.loopVars {
		fmt.Fprintf(g.wr, "int %s_target;\n", s)
	}
	fmt.Fprintf(g.wr, "int target = 0;\n\nsrand((unsigned)time(0));\nfor(;;){switch (target){\ncase 0:\n")
	g.genLines(p.lines)
	fmt.Fprintf(g.wr, "case %d: exit(0);\n", createLabel())
	fmt.Fprintf(g.wr, "default: fprintf(stderr, \"Undefined target %s\", target); abort(); \n}}\n} /* main */\n\n", "%d")

	g.genFunctionDefinitions(p)
	g.genDataDefinitions(p, strType)
	g.genDataDefinitions(p, numType)

	temp := createTemp()
	if temp != 0 {
		fmt.Fprintf(g.wr, "static str temp_str_area[%d], *temp_str = temp_str_area;\n", temp)
	}
}

func (g *generatorForC) genDataDefinitions(p *program, _type astType) {
	size := p.dataCounter[_type]
	fmt.Fprintf(g.wr, "const size_t data_area_for_%s_cnt=%d;\n", _type, size)
	if size == 0 {
		fmt.Fprintf(g.wr, "const %s data_area_for_%s[1]={0};\n\n", _type, _type)
		return
	}
	fmt.Fprintf(g.wr, "const %s data_area_for_%s[%d]={\n", _type, _type, size)
	scan(p, func(h host) {
		c, ok := h.(*cmdData)
		if ok {
			g.generateDataDefinition(c, _type)
		}
	})
	fmt.Fprintf(g.wr, "};\n\n")
}

func (g *generatorForC) genFunctionDeclarations(p *program) {
	b := false
	scan(p, func(h host) {
		switch v := h.(type) {
		case *cmdFnDef:
			g.genDeclaration(v)
			b = true
		}
	})
	if b {
		g.wr.WriteRune('\n')
	}
}

func (g *generatorForC) genFunctionDefinitions(p *program) {
	scan(p, func(h host) {
		switch v := h.(type) {
		case *cmdFnDef:
			g.genDefinition(v)
		}
	})
}

func (g *generatorForC) genPrologue(p *program) {
	g.genFunctionDeclarations(p)
	g.genVarDefinitions(p)
	fmt.Fprintf(g.wr, "static str *temp_str;\n\n")
}

func (g *generatorForC) genVarDefinitions(p *program) {
	m := make(map[string]*astVarRef)
	scan(p, func(h host) {
		switch v := h.(type) {
		case *astVarRef:
			m[v.unambiguousName()] = v
		}
	})
	var names []string
	for k := range m {
		names = append(names, k)
	}
	sort.Strings(names)
	for _, k := range names {
		v := m[k]
		if v.isArray() {
			fmt.Fprintf(g.wr, "static arr %s_var;\n", k)
		} else {
			fmt.Fprintf(g.wr, "static %s %s;\n", v.finalType(), k)
		}
	}
	g.wr.WriteRune('\n')

	for _, k := range names {
		v := m[k]
		if v.isArray() {
			typeString := v.finalType().String()
			fmt.Fprintf(g.wr, "static %s *%s(", typeString, k)
			for i := 0; i != len(v.index); i++ {
				if i != 0 {
					g.wr.WriteRune(',')
				}
				fmt.Fprintf(g.wr, "num index%d", i)
			}
			fmt.Fprintf(g.wr, "){ return %s_in_array(&%s_var,%d", typeString, k, len(v.index))
			for i := 0; i != len(v.index); i++ {
				fmt.Fprintf(g.wr, ",(size_t)index%d", i)
			}
			fmt.Fprintf(g.wr, ");}\n")
		}
	}
	g.wr.WriteRune('\n')
}

func (g *generatorForC) genFunctionHeader(c *cmdFnDef) {
	fmt.Fprintf(g.wr, "static %s fn_%s(", c.expr.finalType(), c.id)
	for i, v := range c.args {
		if i != 0 {
			g.wr.WriteRune(',')
		}
		fmt.Fprintf(g.wr, "num %s", v.unambiguousName())
	}
	fmt.Fprintf(g.wr, ")")
}

func (g *generatorForC) genDeclaration(c *cmdFnDef) {
	g.genFunctionHeader(c)
	fmt.Fprintf(g.wr, ";\n")
}

func (g *generatorForC) genDefinition(c *cmdFnDef) {
	g.genFunctionHeader(c)
	fmt.Fprintf(g.wr, "{\n")
	for _, v := range c.args {
		fmt.Fprintf(g.wr, "\t%s=%s;\n", v.unambiguousName(), v.unambiguousName())
	}
	fmt.Fprintf(g.wr, "\treturn ")
	g.genAstExpr(c.expr)
	fmt.Fprintf(g.wr, ";\n}\n\n")
}

func (g *generatorForC) genAstExpr(a *astExpr) {
	if a.paren {
		g.wr.WriteRune('(')
		g.genAstBoolOp(a.boolOp)
		g.wr.WriteRune(')')
		return
	}
	g.genAstBoolOp(a.boolOp)
}

func (g *generatorForC) genAstBoolOp(a *astBoolOp) {
	g.genAstRelOp(a.head)
	for _, tail := range a.tail {
		g.genAstBoolOpTail(tail)
	}
}

func (g *generatorForC) genAstBoolOpTail(a astBoolOpTail) {
	switch a.oper {
	case tokOr:
		fmt.Fprintf(g.wr, "||")
	case tokAnd:
		fmt.Fprintf(g.wr, "&&")
	}
	g.genAstRelOp(a.val)
}

func (g *generatorForC) genAstRelOp(a *astRelOp) {
	if a.lhs.finalType() == strType && a.rhs != nil {
		g.genAstRelOpForStr(a)
		return
	}
	g.genAstAddOp(a.lhs)
	if a.rhs == nil {
		return
	}
	g.genOperator(a)
	g.genAstAddOp(a.rhs)
}

func (g *generatorForC) genAstRelOpForStr(a *astRelOp) {
	fmt.Fprintf(g.wr, "compare_str(")
	g.genAstAddOp(a.lhs)
	g.wr.WriteRune(',')
	g.genAstAddOp(a.rhs)
	g.wr.WriteRune(')')
	g.genOperator(a)
	g.wr.WriteRune('0')
}

func (g *generatorForC) genOperator(a *astRelOp) {
	switch a.oper {
	case '=':
		fmt.Fprintf(g.wr, "==")
	case tokNe:
		fmt.Fprintf(g.wr, "!=")
	case tokLe:
		fmt.Fprintf(g.wr, "<=")
	case tokGe:
		fmt.Fprintf(g.wr, ">=")
	default:
		fmt.Fprintf(g.wr, "%c", a.oper)
	}
}

func (g *generatorForC) genAstAddOp(a *astAddOp) {
	if a.head.finalType() == strType && len(a.tail) != 0 {
		g.genAspAddOpForStr(a)
		return
	}
	g.genAstMulOp(a.head)
	for _, t := range a.tail {
		g.genAstAddOpTail(t)
	}
}

func (g *generatorForC) genAstAddOpTail(a astAddOpTail) {
	fmt.Fprintf(g.wr, "%c", a.oper)
	g.genAstMulOp(a.val)
}

func (g *generatorForC) genAspAddOpForStr(a *astAddOp) {
	fmt.Fprintf(g.wr, "concat_str(&temp_str[%d],%d,", createTemp(), len(a.tail)+1)
	g.genAstMulOp(a.head)
	for _, t := range a.tail {
		g.wr.WriteRune(',')
		g.genAstMulOp(t.val)
	}
	g.wr.WriteRune(')')
}

func (g *generatorForC) genAstMulOp(a *astMulOp) {
	g.genAstExpOp(a.head)
	for _, t := range a.tail {
		g.genAstMulOpTail(t)
	}
}

func (g *generatorForC) genAstMulOpTail(a astMulOpTail) {
	fmt.Fprintf(g.wr, "%c ", a.oper)
	g.genAstExpOp(a.val)
}

func (g *generatorForC) genAstExpOp(a *astExpOp) {
	if a.rhs == nil {
		g.genAstPart(a.lhs)
		return
	}
	fmt.Fprintf(g.wr, "((num)pow(")
	g.genAstPart(a.lhs)
	g.wr.WriteRune(',')
	g.genAstPart(a.rhs)
	fmt.Fprintf(g.wr, "))")
}

func (g *generatorForC) genAstPart(a *astPart) {
	if a.signal != 0 {
		fmt.Fprintf(g.wr, "%c", a.signal)
	}
	g.genAstValue(a.val)
}

func (g *generatorForC) genAstValue(a astValue) {
	switch v := a.(type) {
	case *astVarRef:
		g.genAstVarRef(v)
	case *astExpr:
		g.genAstExpr(v)
	case *astFnCall:
		g.genAstFnCall(v)
	case *astLit:
		g.genAstLit(v)
	case *astPredef:
		g.genAstPredef(v)
	}
}

func (g *generatorForC) genAstVarRef(a *astVarRef) {
	g.generateCVarRef(a, true)
}

/*
					array        not array     const lit
As an R value
		  str       *s(1)        s             "abc"
		  num       *n(1)        n             1.0000f

As an L value
		  str       let_str(s(1),    let_str(&s,
		  num       let_num(n(1),    let_num(&n,
*/

// generateCVarRef emits C code for access to a variable.
// The intricacies are due to the subtleties of several variable types.
func (g *generatorForC) generateCVarRef(a *astVarRef, shouldDeref bool) {
	if !a.isArray() {
		fmt.Fprintf(g.wr, "%s", a.unambiguousName())
		return
	}
	if shouldDeref {
		g.wr.WriteRune('*')
	}
	fmt.Fprintf(g.wr, "%s", a.unambiguousName())
	g.wr.WriteRune('(')
	for i, v := range a.index {
		if i != 0 {
			g.wr.WriteRune(',')
		}
		g.genAstExpr(v)
	}
	g.wr.WriteRune(')')
}

// gererateCLValue generates code for the left side of an assignment.
// returns a bool indicating wether a closing parenthesis is needed
// The intricacies are due to the subtleties of several variable types.
func (g *generatorForC) generateCLValue(a *astVarRef, fname string) {
	fmt.Fprintf(g.wr, "%s_%s(", fname, a.finalType())
	if !a.isArray() {
		g.wr.WriteRune('&')
	}
	g.generateCVarRef(a, false)
}

func (g *generatorForC) genAstFnCall(a *astFnCall) {
	fmt.Fprintf(g.wr, "fn_%s(", a.id)
	g.genAstExpr(a.arg)
	fmt.Fprintf(g.wr, ")")
}

func (g *generatorForC) genAstLit(a *astLit) {
	switch a._type {
	case numType:
		f, err := strconv.ParseFloat(a.val, 32)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(g.wr, "%.4ff", f)
	case strType:
		g.wr.WriteRune('"')
		for _, r := range a.val {
			if r == '\'' || r == '\\' {
				g.wr.WriteRune('\\')
			}
			g.wr.WriteRune(r)
		}
		g.wr.WriteRune('"')
	}
}

func (g *generatorForC) genAstPredef(a *astPredef) {
	predef := predefs[a.function]
	fmt.Fprintf(g.wr, "%s_%s(", predef.name, predef._type)
	if predef._type == strType {
		fmt.Fprintf(g.wr, "&temp_str[%d],", createTemp())
	}
	for i, arg := range a.args {
		if i != 0 {
			g.wr.WriteRune(',')
		}
		g.genAstExpr(arg)
	}
	g.wr.WriteRune(')')
}

func (g *generatorForC) genLines(pl progLines) {
	for _, l := range pl {
		g.genProgLine(l)
	}
}

func (g *generatorForC) genProgLine(l *progLine) {
	if l.isDst {
		fmt.Fprintf(g.wr, "case %d: ", l.switchID)
	}
	fmt.Fprintf(g.wr, "/* line %d */\n", l.id)
	if len(l.cmds) != 0 {
		g.genCommands(l.cmds)
	}
}

func (g *generatorForC) genCommands(cms cmds) {
	for _, cmd := range cms {
		g.genCmd(cmd)
	}
}

func (g *generatorForC) genCmd(cmd cmd) {
	switch c := cmd.(type) {
	case *cmdData:
	//does nothing. It's not generated inline.
	case *cmdFnDef:
	//does nothing. It's not generated inline.
	case *cmdDim:
		g.genCmdDim(c)
	case *cmdEnd:
		g.genCmdEnd(c)
	case *cmdFor:
		g.genCmdFor(c)
	case *cmdGo:
		g.genCmdGo(c)
	case *cmdIf:
		g.genCmdIf(c)
	case *cmdInput:
		g.genCmdInput(c)
	case *cmdLet:
		g.genCmdLet(c)
	case *cmdNext:
		g.genCmdNext(c)
	case *cmdOn:
		g.genCmdOn(c)
	case *cmdPrint:
		g.genCmdPrint(c)
	case *cmdRead:
		g.genCmdRead(c)
	case *cmdRem:
		g.genCmdRem(c)
	case *cmdRestore:
		g.genCmdRestore(c)
	case *cmdReturn:
		g.genCmdReturn(c)
	case *cmdRun:
		g.genCmdRun(c)
	case *cmdStop:
		g.genCmdStop(c)
	}
}

func (g *generatorForC) generateDataDefinition(c *cmdData, _type astType) {
	emittedSome := false
	//does nothing. Data is not generated here.
	for _, v := range c.values {
		if v.finalType() == _type {
			g.genAstPart(v)
			g.wr.WriteRune(',')
			emittedSome = true
		}
	}
	if emittedSome {
		g.wr.WriteRune('\n')
	}
}

func (g *generatorForC) genCmdDim(c *cmdDim) {
	for _, v := range c.vars {
		fmt.Fprintf(g.wr, "\tdim_%s(&%s_var,%d,", v._type, v.unambiguousName(), len(v.index))
		for i, index := range v.index {
			if i != 0 {
				g.wr.WriteRune(',')
			}
			fmt.Fprintf(g.wr, "(size_t)")
			g.genAstExpr(index)
		}
		fmt.Fprintf(g.wr, ");\n")
	}
}

func (g *generatorForC) genCmdEnd(c *cmdEnd) {
	fmt.Fprintf(g.wr, "\texit(0);\n")
}

func (g *generatorForC) genCmdFor(c *cmdFor) {
	labelInc := createLabel()
	labelCond := createLabel()

	fmt.Fprintf(g.wr, "\t%s_target = %d;\n", c.index.unambiguousName(), labelInc)

	// first index attribution
	g.wr.WriteRune('\t')
	g.generateCLValue(c.index, "let")
	g.wr.WriteRune(',')
	g.genAstExpr(c.begin)
	fmt.Fprintf(g.wr, ");\n")
	fmt.Fprintf(g.wr, "\ttarget = %d; break;\n", labelCond)

	// Increment
	fmt.Fprintf(g.wr, "case %d:\n", labelInc)
	g.wr.WriteRune('\t')
	g.generateCLValue(c.index, "let")
	g.wr.WriteRune(',')
	g.genAstVarRef(c.index)
	g.wr.WriteRune('+')
	g.genAstExpr(c.step)
	fmt.Fprintf(g.wr, ");\n")

	// index value bounds checking
	fmt.Fprintf(g.wr, "case %d:\n", labelCond)
	if c.step == step1 {
		fmt.Fprintf(g.wr, "\tif (")
		g.genAstVarRef(c.index)
		fmt.Fprintf(g.wr, " > ")
		g.genAstExpr(c.end)
		fmt.Fprintf(g.wr, ") { target=%d; break; }\n", c.next.labelExit)
		return
	}
	fmt.Fprintf(g.wr, "\tif (")
	g.genAstExpr(c.step)
	fmt.Fprintf(g.wr, " > 0 && ")
	g.genAstVarRef(c.index)
	fmt.Fprintf(g.wr, " > ")
	g.genAstExpr(c.end)
	fmt.Fprintf(g.wr, ") { target=%d; break; }\n", c.next.labelExit)
	fmt.Fprintf(g.wr, "\telse if (")
	g.genAstExpr(c.step)
	fmt.Fprintf(g.wr, " < 0 && ")
	g.genAstVarRef(c.index)
	fmt.Fprintf(g.wr, " < ")
	g.genAstExpr(c.end)
	fmt.Fprintf(g.wr, ") { target=%d; break; }\n", c.next.labelExit)
}

func (g *generatorForC) genCmdGo(c *cmdGo) {
	if c.sub {
		returnAddress := createLabel()
		fmt.Fprintf(g.wr, "\tpush_address(%d);\n", returnAddress)
		fmt.Fprintf(g.wr, "\ttarget = %d; break;\n", c.dst.adr.switchID)
		fmt.Fprintf(g.wr, "case %d:\n", returnAddress)
	} else {
		fmt.Fprintf(g.wr, "\ttarget = %d; break;\n", c.dst.adr.switchID)
	}
}

func (g *generatorForC) genCmdIf(c *cmdIf) {
	label := c.condBranchTarget()
	if label != 0 {
		g.genConditionalBranch(c, label)
	} else {
		g.genRegularIf(c)
	}
}

func (g *generatorForC) genConditionalBranch(c *cmdIf, label int) {
	fmt.Fprintf(g.wr, "\tif (")
	g.genAstExpr(c.expr)
	fmt.Fprintf(g.wr, "){ target = %d; break; }\n", label)
}

func (g *generatorForC) genRegularIf(c *cmdIf) {
	label := createLabel()
	fmt.Fprintf(g.wr, "\tif (!(")
	g.genAstExpr(c.expr)
	fmt.Fprintf(g.wr, ")){ target = %d; break; }\n", label)
	g.genCommands(c.cmds)
	fmt.Fprintf(g.wr, "case %d:\n", label)
}

func (g *generatorForC) genCmdInput(c *cmdInput) {
	if c.label != "" {
		fmt.Fprintf(g.wr, "\tprint_str(\"%s\");\n", c.label)
	}
	fmt.Fprintf(g.wr, "\tinput_to_buffer();\n")

	for _, v := range c.vars {
		fmt.Fprintf(g.wr, "\t")
		g.generateCLValue(v, "input")
		fmt.Fprintf(g.wr, ");\n")
	}
}

func (g *generatorForC) genCmdLet(c *cmdLet) {
	g.wr.WriteRune('\t')
	g.generateCLValue(c.dst, "let")
	g.wr.WriteRune(',')
	g.genAstExpr(c.src)
	fmt.Fprintf(g.wr, ");\n")
}

func (g *generatorForC) genCmdNext(c *cmdNext) {
	fmt.Fprintf(g.wr, "\ttarget=%s_target; break;\n", c.vars[0].unambiguousName())
	fmt.Fprintf(g.wr, "case %d:\n", c.labelExit)
}

func (g *generatorForC) genCmdOn(c *cmdOn) {
	labelExit := createLabel()
	fmt.Fprintf(g.wr, "\ttarget = (int)(")
	g.genAstExpr(c.expr)
	fmt.Fprintf(g.wr, ");\n")
	fmt.Fprintf(g.wr, "\tif (target < 1 || target > %d) {target = %d; break;}\n", len(c.dsts), labelExit)
	if c.sub {
		fmt.Fprintf(g.wr, "\tpush_address(%d);\n", labelExit)
	}
	fmt.Fprintf(g.wr, "\t{static const int tab[]={")
	for i, line := range c.dsts {
		if i != 0 {
			g.wr.WriteRune(',')
		}
		fmt.Fprintf(g.wr, "%d", line.adr.switchID)
	}
	fmt.Fprintf(g.wr, "}; target = tab[target -1]; break;}\n")
	fmt.Fprintf(g.wr, "case %d:\n", labelExit)
}

func (g *generatorForC) genCmdPrint(c *cmdPrint) {
	g.genPrintSubCmds(c.printSubCmds)
}

func (g *generatorForC) genPrintSubCmds(scs printSubCmds) {
	var _type astType
	shouldNL := true
	for _, subCmd := range scs {
		switch cmd := subCmd.(type) {
		case token:
			if cmd == ';' {
				shouldNL = false
			} else if cmd == ',' {
				shouldNL = false
				fmt.Fprintf(g.wr, "\tprint_char('\\t');\n")
			}
		case *astExpr:
			shouldNL = true
			_type = cmd.finalType()
			if _type == voidType {
				fmt.Fprintf(g.wr, "\t")
				g.genAstExpr(cmd)
				fmt.Fprintf(g.wr, ";\n")
			} else {
				fmt.Fprintf(g.wr, "\tprint_%s(", _type)
				g.genAstExpr(cmd)
				fmt.Fprintf(g.wr, ");\n")
			}
		}
	}
	if shouldNL {
		fmt.Fprintf(g.wr, "\tprint_char('\\n');\n")
	}
}

func (g *generatorForC) genCmdRead(c *cmdRead) {
	for _, v := range c.vars {
		g.wr.WriteRune('\t')
		g.generateCLValue(v, "read")
		fmt.Fprintf(g.wr, ");\n")
	}
}

func (g *generatorForC) genCmdRem(c *cmdRem) {
	fmt.Fprintf(g.wr, "\t/*%s*/\n", c.text)
}

func (g *generatorForC) genCmdRestore(c *cmdRestore) {
	fmt.Fprintf(g.wr, "\trestore();\n")
}

func (g *generatorForC) genCmdReturn(c *cmdReturn) {
	fmt.Fprintf(g.wr, "\tpop_address(&target); break;\n")
}

func (g *generatorForC) genCmdRun(c *cmdRun) {
	fmt.Fprintf(g.wr, "\tsystem(\"%s\");\n", c.arg)
}

func (g *generatorForC) genCmdStop(c *cmdStop) {
	fmt.Fprintf(g.wr, "\texit(0);\n")
}
