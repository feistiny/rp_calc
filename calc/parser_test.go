package calc

import (
	"log"
	"testing"
)

func TestSimpleExp(t *testing.T) {
	exp := "1+2"
	p := Parser{
		exp: exp,
	}
	exp2 := p.exp2Tokens()
	if exp != exp2 {
		log.Fatal("最简单的表达式解析失败")
	}
}

func TestExpWithSpace(t *testing.T) {
	exp := "1 +2"
	expExpected := "1+2"
	p := Parser{
		exp: exp,
	}
	exp2 := p.exp2Tokens()
	if exp2 != expExpected {
		log.Fatal("带空白字符的表达式解析失败")
	}
}

func TestFullExp(t *testing.T) {
	exp := "(1+2*2-4/2)^2"
	p := Parser{
		exp: exp,
	}
	exp2 := p.exp2Tokens()
	if exp2 != exp {
		log.Fatal("带所有支持字符的表达式解析失败")
	}
}
