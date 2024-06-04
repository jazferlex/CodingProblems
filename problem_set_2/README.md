# Problem Set X: Pairs

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
  * Combine the two words
  * Use **isPalindrome** function to determine if the combined words is a palinndrome
  * If it is a palindrome, append it to the `result` list
  * Continue until loop reaches end of the list of words

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
