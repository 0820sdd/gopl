package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/unitconv"
)

func main() {
	flag := 0
	for _, arg := range os.Args[1:] {
		if arg != "" {
			flag = 1
		}
		conv(arg)
	}
	if flag != 1 {
		fmt.Print("请输入数字:\t")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		unitStr := input.Text()
		conv(unitStr)
	}
}

func conv(pa string) {
	t, err := strconv.ParseFloat(pa, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := unitconv.Foot(t)
	m := unitconv.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, unitconv.FToM(f), m, unitconv.MToF(m))
}
