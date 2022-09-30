package codego

import (
	"strings"
)

/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
**/

// func main() {
// 	var long *longest_substring
// 	fmt.Println(long.lengthOfLongestSubString("ertyuioplkjhgfghuioiuytr"))
// 	fmt.Println(long.lengthOfLongestSubString("aaaaaabbbbbbcccccdfghhhj"))
// 	fmt.Println(long.lengthOfLongestSubString("asdfghskl"))
// }

type longest_substring struct{}

func (l *longest_substring) lengthOfLongestSubString(s string) int {
	leng := len(s)
	max := 0
	temStr := ""
	for i := 0; i < leng; i++ {
		ch := s[i]
		index := strings.Index(temStr, string(ch))
		if index >= 0 {
			if index == len(temStr)-1 {
				temStr = string(ch)
			} else {
				temStr = temStr[index+1:] + string(ch)
			}
		} else {
			temStr += string(ch)
		}
		if len(temStr) > max {
			max = len(temStr)
		}
	}
	return max

}
