package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, _ := CountWordsAndImages(os.Args[1])
	fmt.Printf("文字：%d,图片：%d\n", words, images)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

//  练习 5.5： 实现countWordsAndImages。（参考练习4.9如何分词）
func countWordsAndImages(n *html.Node) (words, images int) {
	texts, images := visit(nil, images, n)
	for _, v := range texts {
		v = strings.Trim(strings.TrimSpace(v), "\r\n")
		if v == "" {
			continue
		}
		words += strings.Count(v, "")
	}
	return words, images
}

//递归循环html
func visit(texts []string, images int, n *html.Node) ([]string, int) {
	//文本
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	//图片
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}
		texts, images = visit(texts, images, c)
	}
	//多返回值
	return texts, images
}
