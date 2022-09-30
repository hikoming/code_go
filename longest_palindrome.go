package codego

import (
	"fmt"
)

// func main() {
// 	// var long *longest_palindrome
// 	// fmt.Println(long.longestPalindrome("ertyuioplkjhgfghuioiuytr"))
// 	// fmt.Println(long.longestPalindrome("aaaaaabbbbbbcccccdfghhhj"))
// 	fmt.Println(longestPalindrome("zbarhcagachrbad"))
// }

type longest_palindrome struct{}

// / 回文都是从中心向两边出发，两边相同就是回文，
// / 一个字符串最大的回文，一定是中心点在字符串中间的，次之的中心点往两边偏移
func (l *longest_palindrome) longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	length := len(s)
	center := (length+1)/2 - 1
	ishuiwen := 1
	if length%2 == 0 { //偶数长度
		for i := center; i >= 0; i-- {
			if s[i] != s[length-1-center] { //不是回文
				ishuiwen = 0
				break
			}
		}
	} else { //奇数长度
		for i := center - 1; i >= 0; i-- {
			if s[i] != s[length-1-i] { //不是回文
				ishuiwen = 0
				break
			}
		}
	}
	if ishuiwen == 1 {
		return s
	} else {
		fmt.Println(s[1:])
		sr := l.longestPalindrome(s[1:])
		fmt.Println(s[:length-2])
		sl := l.longestPalindrome(s[:length-1])
		if len(sr) > len(sl) {
			return sr
		} else if len(sr) == len(sl) {
			return sr
		} else {
			return sl
		}
	}
}

//	func (l *longest_palindrome) longest(s string , left int , right int) string {
//		center := (right+left+1)/2 - 1
//		ishuiwen := 1
//		if length%2 == 0 { //偶数长度
//			for i := center; i >= 0; i-- {
//				if s[i] != s[length-1-center] { //不是回文
//					ishuiwen = 0
//					break
//				}
//			}
//		} else { //奇数长度
//			for i := center - 1; i >= 0; i-- {
//				if s[i] != s[length-1-i] { //不是回文
//					ishuiwen = 0
//					break
//				}
//			}
//		}
//		if ishuiwen == 1 {
//			return s
//		} else {
//			fmt.Println(s[1:])
//			sr := l.longestPalindrome(s[1:])
//			fmt.Println(s[:length-2])
//			sl := l.longestPalindrome(s[:length-1])
//			if len(sr) > len(sl) {
//				return sr
//			} else if len(sr) == len(sl) {
//				return sr
//			} else {
//				return sl
//			}
//		}
//	}
//
// //
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
		fmt.Println(s[start : end+1])
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
