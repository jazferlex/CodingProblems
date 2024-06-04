
def isValid(s: str) -> bool:
    prev_length = 0
    while len(s) != prev_length:
        s = s.replace("()", "").replace("[]", "").replace("{}", "")
        prev_length = len(s)
    return s == ""

if __name__ == "__main__":
    test_cases = [
        "()", "()[]{}", "(]", "([)]", "{[]}"
    ]
    for test in test_cases:
        print(isValid(test))