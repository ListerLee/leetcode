package main

import "fmt"

var combinations []string

var m map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backTrack(digits, 0, "")
	return combinations
}

func backTrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := m[digit]
		for _, s := range letters {
			backTrack(digits, index+1, combination+string(s))
		}
	}
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
