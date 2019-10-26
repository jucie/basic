package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type generator interface {
	generate(prog *program)
}

func help() {
	fmt.Println(
		`basic v1.0.1 Copyright(C) 2019 Jucie Dias Andrade
Use:
		basic input.bas output.c`)
}

func main() {
	var outputFormat = flag.String("gen", "c", "output format")
	flag.Parse()

	if len(flag.Args()) < 2 {
		help()
		return
	}

	srcFile, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(flag.Arg(1))
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()

	prog := newProgram()
	rd := bufio.NewReader(srcFile)
	prog.loadFrom(rd)

	wr := bufio.NewWriter(dstFile)
	defer wr.Flush()

	var gen generator
	switch *outputFormat {
	case "c":
		gen = newGeneratorForC(wr)
	default:
		panic(fmt.Sprintf("unknown output format: \"%s\"", *outputFormat))
	}
	gen.generate(prog)
}
