package calc

import (
	"fmt"
	"log"
	"strconv"
)

// 表达式解析器
type Parser struct {
	exp    string   // 原始表达式
	tokens []iToken // 节点: 整数|浮点数|运算符
}

func (p *Parser) SetExp(exp string) {
	p.exp = exp
}

// 解析节点, 词法分析
func (p *Parser) exp2Tokens() (expParsed string) {
	fmt.Println("要解析的表达式:", p.exp)
	multiLen := 0
	for i := 0; i < len(p.exp); i++ {
		curChar := p.exp[i]
		var t iToken

		switch curChar {
		case uint8(OpCharPlus):
			t = NewPlusOperator()
		case uint8(OpCharMinus):
			t = NewMinusOperator()
		case uint8(OpCharMultiple):
			t = NewMultipleOperator()
		case uint8(OpCharDivide):
			t = NewDivideOperator()
		case uint8(OpCharPow):
			t = NewPowOperator()
		case uint8(BracketCharLeft):
			t = NewLeftBracketToken()
		case uint8(BracketCharRight):
			t = NewRightBracketToken()
		case ' ':
			// 空白没意义, 跳过
			println("空白字符跳过不解析")
		default:
			// 目前只支持数字, 其他字符的报错
			if !(curChar >= 0x31 && curChar <= 0x39 || curChar == '.') {
				log.Fatalf("解析到不支持的字符 %s\n", string(curChar))
			}
			multiLen++ // 多字符长度
			if i+1 < len(p.exp) {
				fmt.Printf("解析到多字符的第 %d 个字符 %s\n", multiLen, string(curChar))
				// 不到最后一个字符, 继续循环, 看下一个字符是什么, 如果可以合并就合并,
				// 比如 1.2 就要合并 3 个字符
				continue
			} else {
				// 最后一个字符的时候索引加 1, 保持切分 slice 的时候下标一致
				i++
			}
		}
		if multiLen > 0 {
			mc := p.exp[i-multiLen : i]
			multiLen = 0
			println("解析到多字符的数值", mc)
			f, err := strconv.ParseFloat(mc, 64)
			if err != nil {
				log.Fatal("多字符解析失败", mc, f, err)
			}
			// println("多字符", mc)
			p.tokens = append(p.tokens, NewValueToken(mc))
		}
		if t != nil {
			switch t.(type) {
			case OpToken:
				fmt.Printf("解析到单字符的运算符 %v\n", string(t.Value().(OpChar)))
			case BracketToken:
				fmt.Printf("解析到单字符的括号 %v\n", string(t.Value().(BracketChar)))
			default:
			}
			// 单字符
			p.tokens = append(p.tokens, t)
		}
	}
	for _, token := range p.tokens {
		switch token.(type) {
		case OpToken:
			expParsed += string(token.Value().(OpChar))
		case ValueToken:
			expParsed += token.(ValueToken).ValueString()
		case BracketToken:
			expParsed += string(token.Value().(BracketChar))
		default:
		}
	}
	println("解析后复原表达式:", expParsed)

	return
}
