// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(comma(-5123456.23344))
	fmt.Println(compare("abce", "abde"))
}

//!+
// comma inserts commas in a non-negative decimal integer string.
/*
练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。

练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。
*/
func comma(str float64) string {
	//整型转换成字符串
	s := fmt.Sprintf("%.2f", str)
	//取出小数点后面部分
	var end string
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		end = s[dot:]
		s = s[:dot]
	}
	num := len(s)
	var buf bytes.Buffer
	j := 1
	for i := num - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
		if j%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		j++
	}
	res := buf.String()
	var r bytes.Buffer
	//反转字符串
	for i := len(res) - 1; i >= 0; i-- {
		r.WriteByte(res[i])
	}
	r.WriteString(end)
	return r.String()
}

//练习 3.12： 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
func compare(str1 string, str2 string) bool {
	//比较两个字符串的长度,外层循环是较长的
	num1 := strings.Count(str1, "")
	num2 := strings.Count(str2, "")
	fmt.Println(num1, num2)
	if num2 > num1 {
		str1, str2 = str2, str1
	}
	var res bool
	for _, v := range str1 {
		res = false
		for _, sv := range str2 {
			if v == sv {
				res = true
			}
			fmt.Println(v, sv, res)
		}
		if !res {
			break
		}
	}
	return res
}

//!-
