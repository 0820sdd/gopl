package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	outline(os.Args[1])
}

func outline(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	doc, _ := html.Parse(resp.Body)
	//1.使用函数值
	// forEachNode(doc, startElement, endElement)
	//2.
	doc2 := ElementByID(doc, "go-import", startElement2)
	forEachNode(doc2, startElement, endElement)
	resp.Body.Close()

	s := expand("footest", expand2)
	fmt.Println(s)
	return "", nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	//显式的调用一下
	if pre != nil {
		pre(n)
	}

	//fmt.Println(n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

/*
练习 5.7： 完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。使用简略格式输出没有孩子结点的元素（即用<img/>代替<img>
</img>）。编写测试，验证程序输出的格式正确。（详见11章）
优化了script标签问题
*/
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attr := ""
		for _, a := range n.Attr {
			attr += " " + a.Key + "=" + "\"" + a.Val + "\" "
		}
		fmt.Printf("%*s<%s%s", depth*2, "", n.Data, attr)
		depth++
	}
	if n.Type == html.ElementNode && n.FirstChild == nil && n.Data != "script" {
		fmt.Printf("/>\n")
	} else if n.Type == html.ElementNode {
		fmt.Printf(">\n")
	}

	if n.Type == html.TextNode {
		fmt.Printf("%*s %s\n", depth*2, "", n.Data)
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild == nil && n.Data != "script" {
		depth--
		fmt.Printf("\n")
		return
	}
	if n.Type == html.ElementNode {
		depth--

		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

/*
练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNode的遍历。使用修改后的代码编写ElementByID函数，
根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。

func ElementByID(doc *html.Node, id string) *html.Node
*/
func ElementByID(n *html.Node, idStr string, pre func(*html.Node, string) bool) *html.Node {
	//显式的调用一下
	if pre != nil {
		if pre(n, idStr) {
			return n
		}
	}

	//fmt.Println(n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ElementByID(c, idStr, pre)
	}
	return n
}

func startElement2(n *html.Node, idStr string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == idStr {
				fmt.Println("==========", a.Val)
				break
				return true

			}
		}
	}
	return false
}

/*
练习 5.9： 编写函数expand，将s中的"foo"替换为f("foo")的返回值。
*/
func expand(s string, f func(string) string) string {
	str := f("foo")
	s = strings.Replace(s, "foo", str, -1)
	return s
}
func expand2(s string) string {
	return s + "-test"
}
