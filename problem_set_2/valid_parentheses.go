package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	prev_length := 0
	for len(s) != prev_length{
		s = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s,"()",""),"[]",""),"{}","")
		prev_length = len(s)
	}
	return s == ""
}

func main() {
	test_cases := []string{
        "()", "()[]{}", "(]", "([)]", "{[]}",
	}
	for _, test := range test_cases{
		result := isValid(test)
		fmt.Println(result)
	}
}
	