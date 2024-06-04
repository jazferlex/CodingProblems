# Problem Set X: Palindrome Pairs

## Problem Description

    Given a list of words as input, find pairs of indices that can make these two words into a palindrome.

## Solution Overview

**isPalindrome Function**

Return true if the word is equal to its reverse.

* **Python:**
  * Python has a very straightforward way to reverse strings and arrays with `[::-1]`
* **Go:**
  * Using a reverse for loop to go through the word and appending each letter to a new string

**palindromePairs Function**

Return a list of pairs of indices that can make these two words into a palindrome.

* **Python and Go:**
  * Using two pointers to iterate through the array of words
  * To avoid pointing to the same word, `pointer2` should always be greater than `pointer1`
  * Combine the two words
  * Use **isPalindrome** function to determine if the concatenation of these words is a palindrome
  * If it is a palindrome, append the pair of indices to the `result` list

## Instructions to Run the Code

Assuming the terminal is already inside **problem_set_1** folder

* **Python:**

  ```
  py palindrome_pairs.py
  ```
* **Go:**

  ```
  go run palindrome_pairs.go
  ```
