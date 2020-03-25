// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	// var counts = make(map[string]int)
	// counts = count(counts, doc)
	// for k, v := range counts {
	for k, v := range count(map[string]int{}, doc) {
		fmt.Printf("%s ==> %d\n", k, v)
	}

	for _, texts := range visit2(nil, doc) {
		fmt.Println(texts)
	}

	for _, res := range visit3(nil, doc) {
		fmt.Println(res)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	// 练习 5.1： 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
	// 实在是不知道为啥不对，我选择放弃
	// if c := n.FirstChild; c != nil {
	// 	links = visit(links, c)
	// } else if n.NextSibling != nil {
	// 	links = visit(links, n.NextSibling)
	// }

	return links
}

//!-visit

// 练习 5.2： 编写函数，记录在HTML树中出现的同名元素的次数。
func count(counts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		counts = count(counts, c)
	}
	return counts
}

// 练习 5.3： 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。
func visit2(texts []string, n *html.Node) []string {
	// if n.Type == html.TextNode && n.Data != "script" && n.Data != "style" {
	// 	texts = append(texts, n.Data)
	// }
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	texts = visit2(texts, c)
	// }
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}
		texts = visit2(texts, c)
	}
	return texts
}

/*
练习 5.4： 扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。
*/
func visit3(res []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "img" || n.Data == "link" || n.Data == "scripts") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				res = append(res, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = visit3(res, c)
	}
	return res
}
