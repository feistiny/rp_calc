package calc

import (
	"log"
	"testing"
)

func TestSimpleRP(t *testing.T) {
	c := NewCalc("1+1")
	res := c.Result()
	if c.rp != "1 1 +" {
		log.Fatal("简单的逆波兰表达式转换失败: ", c.rp)
	}
	if res != float64(2) {
		log.Fatal("简单的逆波兰表达式求值失败: ", res)
	}
}

func TestZhiHuRP(t *testing.T) {
	c := NewCalc("1+((2+3)*4)-5")
	res := c.Result()
	if c.rp != "1 2 3 + 4 * + 5 -" {
		log.Fatal("知乎案例的逆波兰表达式转换失败: ", c.rp)
	}
	if res != float64(16) {
		log.Fatal("知乎案例的逆波兰表达式求值失败: ", res)
	}
}

func TestFullRP(t *testing.T) {
	c := NewCalc("(4/2+2^3-1)*2")
	res := c.Result()
	if c.rp != "4 2 / 2 3 ^ + 1 - 2 *" {
		log.Fatal("全运算符的逆波兰表达式转换失败: ", c.rp)
	}
	if res != float64(18) {
		log.Fatal("全运算符的逆波兰表达式转换失败: ", res)
	}
}
