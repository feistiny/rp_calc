package main

import (
	"feistiny/calc"
	"fmt"
	"strconv"
)

func main() {
	var exp string
	println("请输入要计算的表达式(支持加、减、乘、除、幂运算和括号): ")
	_, _ = fmt.Scan(&exp)

	c := calc.NewCalc(exp)
	res := c.Result()
	println("计算的最终结果是:", strconv.FormatFloat(res, 'f', -1, 64))
}
