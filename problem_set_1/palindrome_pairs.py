
def isPalindrome(word):
    return word == word[::-1]

def palindromePairs(words):
    result = []
    for pointer1, word1 in enumerate(words):
        for pointer2, word2 in enumerate(words):
            if isPalindrome(word1+word2):
                result.append([pointer1, pointer2])
    return result

test_cases = [
    ["bat", "tab", "cat"],
    ["dog", "cow", "god", "woc"],
    ["cat", "dog", "god", "tcc"],
    ["cat", "dog", "ogd", "act", "dat"]
]

if __name__ == '__main__':
    for words in test_cases:
        print(palindromePairs(words))
