package main

import (
	"fmt"

	tempconv "gopl.io/ch2/2.1"
)

func main() {

	//练习 2.1： 向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。
	//显式类型转换
	a := tempconv.Celsius(tempconv.AbsoluteZeroK)

	b := tempconv.CToK(a)
	fmt.Println(b)
	c := tempconv.KToC(100)
	fmt.Println(c)
}
