package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	for prev_length:=0; prev_length != len(s); prev_length = len(s){
		s = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s,"()",""),"[]",""),"{}","")
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
	