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
	words []uint64
}

func main() {

	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	x.AddAll(1, 9, 144)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	// x.UnionWith(&y)
	// fmt.Println(x.String()) // "{1 9 42 144}"

	z := x.IntersectWith(&y)
	fmt.Println(z.String())

	a := x.DifferenceWith(&y)
	fmt.Println(a.String())

	b := x.SymmetricDifferenceWith(&y)
	fmt.Println(b.String())

	c := x.Elems()
	fmt.Println(c)
}

// 练习 6.2： 定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)。

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
	fmt.Println("words is", s.words)
}

// AddAll adds the non-negative value x to the set.
// func (s *IntSet) AddAll(elements ...int) {
// 	for _, x := range elements {
// 		word, bit := x/64, uint(x%64)
// 		for word >= len(s.words) {
// 			s.words = append(s.words, 0)
// 		}
// 		s.words[word] |= 1 << bit
// 	}
// }

// AddAll adds the non-negative value x to the set.
func (s *IntSet) AddAll(elements ...int) {
	for _, x := range elements {
		s.Add(x)
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

//!+UnionWith

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

//!-UnionWith

// 练习 6.3： (*IntSet).UnionWith会用|操作符计算两个集合的交集，我们再为IntSet实现另外的几个函数IntersectWith(交集：元素在A集合B集合均出现),
// DifferenceWith(差集：元素出现在A集合，未出现在B集合),SymmetricDifference(并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A)。

//!+IntersectWith

// IntersectWith sets s to the Intersect of s and t.
func (s *IntSet) IntersectWith(t *IntSet) IntSet {
	var result IntSet
	for i, word := range s.words {
		if i >= len(t.words) {
			break
		}
		result.words = append(result.words, word&t.words[i])
	}
	return result
}

//!+DifferenceWith

// DifferenceWith sets s to the Difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) IntSet {
	var result IntSet
	for i, word := range s.words {
		if i >= len(t.words) {
			result.words = append(result.words, word)
			continue
		}
		result.words = append(result.words, word&(word^t.words[i]))
	}
	return result
}

//!-DifferenceWith

//!+SymmetricDifferenceWith

// SymmetricDifferenceWith sets s to the SymmetricDifference of s and t.
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) IntSet {
	var result IntSet
	for i, word := range s.words {
		if i >= len(t.words) {
			result.words = append(result.words, word)
			continue
		}
		result.words = append(result.words, word^t.words[i])
	}
	return result
}

//!-SymmetricDifferenceWith

// 练习6.4: 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。
//!+Elems

// Elems returns 。。。
func (s *IntSet) Elems() []int {
	var result []int
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, 64*i+j)
			}
		}
		fmt.Println(",,,,,,", result)
	}
	return result

}

//!-Elems
