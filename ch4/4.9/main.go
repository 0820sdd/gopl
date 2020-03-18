// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
*/

func main() {
	seen := make(map[string]int) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	// for input.Scan() {
	// 	line := input.Text()
	// 	seen[line]++
	// }
	// fmt.Printf("words\tcount\n")
	// for c, n := range seen {
	// 	fmt.Printf("%q\t%d\n", c, n)
	// }
	for input.Scan() {
		seen[input.Text()]++
	}
	for k, v := range seen {
		fmt.Printf("%s == %d \n", k, v)
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

//!-
