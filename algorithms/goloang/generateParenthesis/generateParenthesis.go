package main

import "fmt"

func generateParenthesis1(n int) []string {
	cur := make([]byte, 2*n)
	ans := []string{}
	generateAll(cur, 0, &ans)
	return ans
}

func generateAll(current []byte, pos int, ans *[]string) {

	if len(current) == pos {
		if isValid(current) {
			*ans = append(*ans, string(current))
		}

	} else {
		current[pos] = '('
		generateAll(current, pos+1, ans)
		current[pos] = ')'
		generateAll(current, pos+1, ans)
	}
}

func isValid(current []byte) bool {
	n := len(current)
	var balance int = 0
	if n%2 != 0 {
		return false
	}

	for _, v := range current {
		if v == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

func generateParenthesis(n int) []string {
	cur := make([]byte, 2*n)
	ans := []string{}
	backTrace(cur, 0, &ans, 0, 0, n)
	return ans
}

func backTrace(cur []byte, pos int, ans *[]string, open, close, max int) {
	if len(cur) == pos {
		*ans = append(*ans, string(cur))
		return
	}
	if open < max {
		cur[pos] = '('
		backTrace(cur, pos+1, ans, open+1, close, max)
	}
	if close < open {
		cur[pos] = ')'
		backTrace(cur, pos+1, ans, open, close+1, max)
	}
}

func main() {
	n := 2
	ans := generateParenthesis(n)
	fmt.Println(ans)
}
