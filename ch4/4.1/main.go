package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

//命令行标志
var hashMethod = flag.String("s", "sha256", "请输入哈希算法")

func main() {
	flag.Parse()
	num := compareSha256("x", "X")
	fmt.Println(num)

	//接收命令行flag,进行判断
	//printHash(strings.ToUpper(*hashMethod), "x")
	printHash(*hashMethod, "x")
}

/*
练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
*/
func compareSha256(str1 string, str2 string) int {
	a := sha256.Sum256([]byte(str1))
	b := sha256.Sum256([]byte(str2))
	num := 0
	//循环字节数组
	for i := 0; i < len(a); i++ {
		//1个字节8个bit,移位运算，获取每个bit
		for m := 1; m <= 8; m++ {
			//比较每个bit是否相同
			if (a[i] >> uint(m)) != (b[i] >> uint(m)) {
				num++
			}
		}
	}
	return num
}

/*
练习 4.2： 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
*/
func printHash(flag string, str string) {
	if flag == "sha256" {
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
		return
	}
	if flag == "sha512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
		return
	}
	if flag == "sha384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(str)))
		return
	}

}
