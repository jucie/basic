package main

import (
	"bufio"
	"fmt"
	"os"
)

var source []string

func loadSourceFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		source = append(source, line)
	}
}

func generateExecutable() {
	for _, line := range source {
		fmt.Println(line)
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Missing args")
	}
	for i := 1; i < len(os.Args); i++ {
		path := os.Args[i]
		loadSourceFile(path)
		generateExecutable()
	}
}
