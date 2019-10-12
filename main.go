package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type generator interface {
	generate(wr *bufio.Writer, prog *program)
}

func main() {
	var outputFormat = flag.String("gen", "c", "output format")
	var srcFilePath = flag.String("in", "", "input file")
	var dstFilePath = flag.String("out", "", "output file")
	flag.Parse()

	var err error
	var srcFile = os.Stdin
	var dstFile = os.Stdout

	if *srcFilePath != "" {
		srcFile, err = os.Open(*srcFilePath)
		if err != nil {
			panic(err)
		}
		defer srcFile.Close()
	}
	if *dstFilePath != "" {
		dstFile, err = os.Create(*dstFilePath)
		if err != nil {
			panic(err)
		}
		defer dstFile.Close()
	}

	prog := newProgram()
	rd := bufio.NewReader(srcFile)
	prog.loadFrom(rd)

	var gen generator
	switch *outputFormat {
	case "c":
		gen = newGeneratorForC()
	default:
		panic(fmt.Sprintf("unknown output format: \"%s\"", *outputFormat))
	}
	wr := bufio.NewWriter(dstFile)
	defer wr.Flush()
	gen.generate(wr, prog)
}
