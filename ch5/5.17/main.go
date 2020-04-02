package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, _ := http.Get("http://mail.sina.net")
	doc, _ := html.Parse(resp.Body)
	resp.Body.Close()
	tNodes := ElementsByTagName(doc, "p", "a")
	fmt.Println("=========", &tNodes)
	for _, v := range tNodes {
		ForEachNode(v)
		fmt.Println("------------------------")
	}
}

var nodes []*html.Node

/*
练习5.17：编写多参数版本的 ElementsByTagName ，函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。下面给出了2个例子：
// func ElementsByTagName(doc *html.Node, name...string) []*html.Node
// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
*/
// ElementsByTagName is a func
func ElementsByTagName(n *html.Node, names ...string) []*html.Node {
	// newNode := n
	for _, name := range names {
		if n.Type == html.ElementNode && n.Data == name {
			nodes = append(nodes, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//可变参数传参特点
		ElementsByTagName(c, names...)
	}
	return nodes
}

//ForEachNode is a func
func ForEachNode(n *html.Node) {
	fmt.Println(n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c)
	}
}
