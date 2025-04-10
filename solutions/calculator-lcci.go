package solutions

import "strings"

// 给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
// 表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
// 示例 1:
// 输入: "3+2*2"
// 输出: 7
// 示例 2:
// 输入: " 3/2 "
// 输出: 1
// 示例 3:
// 输入: " 3+5 / 2 "
// 输出: 5
// 说明：
// 你可以假设所给定的表达式都是有效的。
// 请不要使用内置的库函数 eval。

// Calculate 百度二面
func Calculate(s string) int {
	var stack []int
	s = strings.Replace(s, " ", "", -1)
	var operator = '+'
	var num = 0
	for i, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}
		if c < '0' || c > '9' || i == len(s)-1 {
			switch operator {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			num = 0
			operator = c
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}
