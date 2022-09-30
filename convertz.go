package codego

// func main() {
// 	var c *convertz
// 	// fmt.Println(long.lengthOfLongestSubString("ertyuioplkjhgfghuioiuytr"))
// 	// fmt.Println(long.lengthOfLongestSubString("aaaaaabbbbbbcccccdfghhhj"))
// 	fmt.Println(c.convert("asdfghskl"))
// }

type convertz struct{}

func (c *convertz) convert(s string, numRows int) {
	// index := 0
	str := ""
	for i := 0; i < numRows; i++ {
		jiange := numRows*2 - 1
		n := len(s) / (jiange - 1)
		if len(s)%(jiange-1) > 0 {
			n++
		}
		for i := 0; i < n; i++ {
			str += string(s[i])
		}

	}
}
