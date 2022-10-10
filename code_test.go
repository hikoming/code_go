package codego

import (
	"fmt"
	"sync"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Printf("isPalindrome:%t", isPalindrome("abcbA"))
	fmt.Println("")
}

func TestPartition(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	var wg sync.WaitGroup
	for k, v := range m {
		wg.Add(1)
		go func(k int, v int, wg *sync.WaitGroup) {
			fmt.Println("k:", k, "v:", v)
			wg.Done()
		}(k, v, &wg)
	}
	wg.Wait()
}
