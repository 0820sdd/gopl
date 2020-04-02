// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import "fmt"

// 练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
//!+
func max(vals ...int) int {
	maxx := 0
	for _, val := range vals {
		if val > maxx {
			maxx = val
		}
	}
	return maxx
}

func min(vals ...int) (int, error) {
	var minn int
	lens := len(vals)
	if lens == 0 {
		return 0, fmt.Errorf("max: %s", "至少传递一个参数")
	}
	for _, val := range vals {
		if val < minn {
			minn = val
		}
	}
	return minn, nil
}

//!-

// 练习5.16：编写多参数版本的strings.Join。
func join(sep string, strs ...string) (string, error) {

	var res string
	for k, v := range strs {
		if k == (len(strs) - 1) {
			res = res + v
		} else {
			res += v + sep
		}

	}
	return res, nil
	// strings包中Join的实现
	// n := len(sep) * (len(strs) - 1)
	// for i := 0; i < len(strs); i++ {
	// 	n += len(strs[i])
	// 	fmt.Println("n is:", n)
	// }

	// b := make([]byte, n)
	// bp := copy(b, strs[0])
	// fmt.Println("----", bp)
	// for _, s := range strs[1:] {
	// 	bp += copy(b[bp:], sep)
	// 	bp += copy(b[bp:], s)
	// }
	// return string(b), nil
}

func main() {
	//!+main
	fmt.Println(max())                      //  "0"
	fmt.Println(max(3))                     //  "3"
	fmt.Println("max is:", max(1, 2, 3, 4)) //
	//!-main
	fmt.Println(min(1, 4, 3, 4, 5)) //

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println("max is:", max(values...)) //
	fmt.Println(min(values...))            //
	//!-slice
	fmt.Println(join("-", "hello", "world", "and"))
}
