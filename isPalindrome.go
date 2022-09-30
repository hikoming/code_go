package codego

func isPalindrome(s string) bool {
	var sgood string
	n := len(s)
	for i := 0; i < n; i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}
	n = len(sgood)
	for i := 0; i < n/2; i++ {
		if sgood[i] != sgood[n-i-1] {
			return false
		}
	}
	return true

}

func isalnum(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	} else {
		return false
	}
}
