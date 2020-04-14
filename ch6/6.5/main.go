// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

const (
	bitNum = (32 << (^uint(0) >> 63)) //根据平台自动判断决定是32还是64
)

func main() {

	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	x.Remove(9)
	fmt.Println(x.String()) // "{1 9 144}"

	z := x.Copy()
	fmt.Println(z.String())
	z.Clear()
	fmt.Println(z.String())

}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bitNum, uint(x%bitNum)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bitNum, uint(x%bitNum)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitNum*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

// 练习6.1: 为bit数组实现下面这些方法

// func (*IntSet) Len() int      // return the number of elements
// func (*IntSet) Remove(x int)  // remove x from the set
// func (*IntSet) Clear()        // remove all elements from the set
// func (*IntSet) Copy() *IntSet // return a copy of the set

//!+len

// Len returns the lens
func (s *IntSet) Len() int {
	lens := 0
	for _, word := range s.words {
		for j := 0; j < bitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				lens++
			}
		}
	}
	return lens
}

//!-len

//!+Remove

// Remove remove the non-negative value x from the set.
//删除集合中的元素
//1.异或^ ：两个值相同，结果为0；两个值不同结果为1；
//2.与&：两个值都是1，结果为1；其他结果为0
//3. s.words[word] ^ (1 << bit) 把我指定位的1改成了0
//4. a &= b  ==>  a=a&b  最终实现设置指定位为0

func (s *IntSet) Remove(x int) {
	word, bit := x/bitNum, uint(x%bitNum)
	s.words[word] &= s.words[word] ^ (1 << bit)
}

//!-Remove

//!+Clear

// Clear remove the non-negative value x from the set.
//清空集合
//1. 设置每个位都为0
//2. 使用异或,把位是1的改成0

func (s *IntSet) Clear() {
	for i, word := range s.words {
		for j := 0; j < bitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				s.words[i] ^= 1 << uint(j)
			}
		}
	}
}

//!-Copy

// Copy remove the non-negative value x from the set.
func (s *IntSet) Copy() (r *IntSet) {
	var results IntSet
	for _, word := range s.words {
		results.words = append(results.words, word)
	}
	return &results
}

//!-Copy
