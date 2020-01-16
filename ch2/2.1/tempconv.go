// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+
// 练习 2.1： 向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Kelvin float64

const (
	AbsoluteZeroK Kelvin = -273.15
	FreezingK     Kelvin = 0
	BoilingK      Kelvin = 100
)

func (k Kelvin) String() string  { return fmt.Sprintf("%g°K", k) }
func (c Celsius) String() string { return fmt.Sprintf("%g °c", c) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c) }

//!-
