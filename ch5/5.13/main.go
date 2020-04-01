package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"gopl.io/ch5/links"
)

/*
练习5.13： 修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。只保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。
*/
var sum int

func main() {
	breadthFirst(crawl, os.Args[1:])
}

/*
抓取页面的所有连接
*/
func crawl(url string) []string {
	sum++

	go save(url)
	fmt.Printf("%d|%s\n", sum, url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

/*
保存页面到文件
*/
func save(u string) bool {

	urlObj, _ := url.Parse(u)
	path := "/tmp/crawl/" + urlObj.Host
	if urlObj.Path == "" || urlObj.Path == "/" {
		urlObj.Path = "/index.html"
	}
	filename := path + urlObj.Path //重点注意文件名
	fmt.Println(filename)
	//打开文件
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	fmt.Println("f is:", f)
	//读取链接
	resp, geterr := http.Get(u)

	if geterr != nil || resp.StatusCode != http.StatusOK {
		//resp.Body.Close()
		return false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	//创建保存目录
	_, err := os.Stat(path)
	if err != nil {
		os.MkdirAll(path, 0755)
	}

	io.WriteString(f, string(body))
	resp.Body.Close()
	body = nil
	return true
}

/*
广度优先算法
*/
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
