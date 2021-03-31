package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	if n%2 == 1 {
		return false
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if m[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	s := "()"
	fmt.Println(isValid(s))
}
