package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	a := PopCount1(25)
	fmt.Println(a)
	n := PopCountFor(80)
	fmt.Println(n)
}

//练习 2.4： 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。
//移位算法
func PopCount1(x uint64) int {
	num := 0
	for i := 0; x != 0; x = x >> 1 {
		if x&1 == 1 {
			i++
		}
		num = i
	}
	return num
}

//练习 2.5： 表达式x&(x-1)用于将x的最低的一个非零的bit位清零。使用这个算法重写PopCount函数，然后比较性能。
//x & (x-1)算法
func PopCount2(x uint64) int {
	num := 0
	for x != 0 {
		x = x & (x - 1)
		num++
	}
	return num
}

//查表法
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//练习 2.3： 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。（11.4节将展示如何系统地比较两个不同实现的性能。）
//查表法,使用循环
func PopCountFor(x uint64) int {
	var num byte
	var i uint64
	for i = 0; i < 8; i++ {
		num += pc[byte(x>>(i*uint64(8)))]
	}
	return int(num)
}
