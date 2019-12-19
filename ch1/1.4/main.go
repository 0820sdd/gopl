
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
// 练习 1.4： 修改dup2，出现重复的行时打印文件名称。
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type lines struct {
	FileName string
	String   string
}

func main() {
	counts := make(map[lines]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "none")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, f.Name())
			f.Close()
		}
	}
	for lines, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, lines.String, lines.FileName)
		}
	}
}

func countLines(f *os.File, counts map[lines]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[lines{
			FileName: filename,
			String: input.Text(),
		}]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
