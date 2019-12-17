// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
// 练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println( i, os.Args[i+1])
	}
	fmt.Println(s, os.Args[0])
}

//!-
