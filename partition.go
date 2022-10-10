package codego

import "fmt"

// func partition(s string) [][]string {
func partition1() {
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			ch1 <- i
		} else {
			ch2 <- i
		}
	}
	for {
		select {
		case i := <-ch1:
			fmt.Println(i)
		case j := <-ch2:
			fmt.Println(j)
		default:
			fmt.Println("")
			return
		}
	}
}

func partition(s string) [][]string {
	var res [][]string
	var path []string
	size := len(s)
	var backTrack func(int)
	backTrack = func(start int) {
		// 递归终止条件
		// 如果切割线，也就是start移动到了字符串s末尾之后，说明我们找到了一个符合要求的切割方案
		if start == size {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := start; i < size; i++ {
			// 如果s[start:i+1]不是回文子串，就可以跳过本次循环了
			if Check(s[start : i+1]) {
				path = append(path, s[start:i+1])
			} else {
				continue
			}
			// 递归
			backTrack(i + 1)
			// 回溯
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}

func Check(s string) bool {
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
