package main

import "fmt"

func isPalindrome(s string) bool {
    reverse := ""
    for i := len(s) - 1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return s == reverse
}

func palindromePairs(words []string) [][]int {
    result := [][]int{}
    for pointer1, word1 := range words {
        for pointer2, word2 := range words {
            if isPalindrome(word1 + word2) {
                result = append(result, []int{pointer1, pointer2})
            }
        }
    }
    return result
}

func main() {
    test_cases := [][]string{
        {"bat", "tab", "cat"},
        {"dog", "cow", "god", "woc"},
        {"cat", "dog", "god", "tcc"},
        {"cat", "dog", "ogd", "act", "dat"},
    }
    for _, inputWords := range test_cases {
        result := palindromePairs(inputWords)
        fmt.Println(result)
    }
}
