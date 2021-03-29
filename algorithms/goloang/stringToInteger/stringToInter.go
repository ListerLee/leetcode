package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {

	status := "start"
	var signed int = 1
	var ans int64
	m := map[string][]string{
		"start":  {"start", "signed", "count", "stop"},
		"signed": {"stop", "stop", "count", "stop"},
		"count":  {"stop", "stop", "count", "stop"},
		"stop":   {"stop", "stop", "stop", "stop"},
	}

	getStatus := func(c byte) int {
		if c == ' ' {
			return 0
		} else if c == '+' || c == '-' {
			return 1
		} else if c >= '0' && c <= '9' {
			return 2
		} else {
			return 3
		}
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		status = m[status][getStatus(c)]

		if status == "signed" {
			if c == '-' {
				signed = -1
			} else {
				signed = 1
			}
		} else if status == "count" {
			ans = ans*10 + int64(c-0x30)
			if signed == 1 {
				ans = min(ans, int64(math.MaxInt32))
			} else {
				ans = min(ans, -int64(math.MinInt32))
			}
		} else if status == "stop" {
			break
		}
	}
	return int(ans) * signed
}

func min(x, y int64) int64 {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	str := "-91283472332"

	fmt.Println(myAtoi(str))

}
