// This program tests to find out if all files in list are BASIC old school.
package main

import (
	"bufio"
	"os"
	"sync"
	"unicode"
)

func evalueteFile(wg *sync.WaitGroup, path string) {
	defer wg.Done()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if len(line) > 0 && !unicode.IsDigit(rune(line[0])) {
			println(path, " no es")
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Missing list file")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var wg sync.WaitGroup
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		wg.Add(1)
		go evalueteFile(&wg, line)
	}
	wg.Wait()
}
