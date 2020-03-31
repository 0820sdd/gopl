// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	for key, value := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", key, value)
	}

}

// 练习5.10： 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。
func topoSort(m map[string][]string) map[int]string {
	var order = make(map[int]string)
	index := 1
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// order = append(order, item)
				visitAll(m[item])
				order[index] = item
				index++
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	// sort.Strings(keys)
	visitAll(keys)
	return order
}

//!-main

/*
练习5.11： 现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图中的环。
等着去看数据结构再看这个题
*/
