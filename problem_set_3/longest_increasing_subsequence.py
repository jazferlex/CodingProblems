def lengthOfLIS(nums):
    result = [nums[0]]
    for num in nums[1:]:
        if num > result[-1]:
            result.append(num)
        else:
            for i, n in enumerate(result):
                if num <= n:
                    result[i] = num
                    break
    return len(result)

if __name__ == '__main__':
    test_cases = [
        [10, 9, 2, 5, 3, 7, 101, 18],  # 4
        [0, 1, 0, 3, 2, 3],  # 4
        [7, 7, 7, 7, 7, 7, 7],  # 1
        [4, 10, 4, 3, 8, 9],  # 3
        [2, 2],  # 1
        [1, 3, 6, 7, 9, 4, 10, 5, 6],  # 6
    ]
    for test in test_cases:
        print(lengthOfLIS(test))
