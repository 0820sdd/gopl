// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	//!+array
	a := [6]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	b := []int{0, 1, 2, 3, 4, 5}
	br := rotate(b, 2)
	fmt.Println(br)
	c := []string{"test", "fd", "hello", "hello", "world"}
	cr := emptyString(c)
	fmt.Println(cr)
	d := []byte("abc bcd wer  sdsd  hello     de")
	dr := emptyString2(d)
	fmt.Println(string(dr))

	e := []byte("abc bcd wer  sdsd  world     de")
	reverse2(e)
	fmt.Println(string(e))

	//!-array

	// Interactive test of reverse.
	// 	input := bufio.NewScanner(os.Stdin)
	// outer:
	// 	for input.Scan() {
	// 		var ints [6]int
	// 		for i, s := range strings.Fields(input.Text()) {

	// 			x, err := strconv.ParseInt(s, 10, 64)
	// 			if err != nil {
	// 				fmt.Fprintln(os.Stderr, err)
	// 				continue outer
	// 			}
	// 			ints[i] = int(x)
	// 			fmt.Println(x)
	// 		}
	// 		reverse(&ints)
	// 		fmt.Printf("%v\n", ints)
	// 	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
/*
练习 4.3： 重写reverse函数，使用数组指针代替slice。
*/
func reverse(s *[6]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

/*
练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
*/
func rotate(x []int, y int) []int {
	lens := len(x)
	res := make([]int, lens)
	for i := range x {
		num := i + y
		// if num < len(x) {
		// 	res[i] = x[num]
		// } else {
		// 	res[i] = x[num-len(x)]
		// }
		if num >= lens {
			num = num - lens
		}
		res[i] = x[num]

	}
	return res
}

/*
练习 4.5：写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
*/
func emptyString(x []string) []string {

	lens := len(x)
	res := make([]string, lens)
	index := 0
	i := 0
	for _, v := range x {
		index = i + 1
		if index >= lens {
			break
		}
		if v != x[index] {
			res[i] = v
			i++
		}
	}
	return res[:i]
}

/*
练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
*/
func emptyString2(x []byte) []byte {
	lens := len(x)
	index := 0
	for i, v := range x {
		index = i + 1
		if index >= lens {
			break
		}
		if unicode.IsSpace(rune(v)) && unicode.IsSpace(rune(x[index])) {
			copy(x[i:], x[index:])
			x = x[:len(x)-1]
			i--
		}
	}
	return x
}

/*
练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
*/
func reverse2(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
