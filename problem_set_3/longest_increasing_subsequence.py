def length_of_lis(nums):
    result = [nums[0]]
    for num in nums[1:]:
        if num > result[-1]:
            result.append(num)
        else:
            i = 0
            while i < len(result) and result[i] < num:
                i += 1
            result[i] = num
    return len(result)

if __name__ == '__main__':
    test_cases = [
        [10, 9, 2, 5, 3, 7, 101, 18],
        [0, 1, 0, 3, 2, 3],
        [7, 7, 7, 7, 7, 7, 7],
        [4, 10, 4, 3, 8, 9]
    ]
    for test in test_cases:
        print(length_of_lis(test))
