# Problem Explanation
The idea is to find the maximum frequency of each letter in words2, and then check if each word in words1 has at least the same frequency of each letter.<br>
For example, if words2 is ["el", "llo"], the maximum frequency of each latter is<br>
{'e': 1, 'l': 2}.<br>

Then, we check if each word in words1 has the frequency of each letter is greater than or equal to the maximum frequency of each letter in words2.

Let's summarize the steps to solve the problem:
1. Find the maximum frequency of each letter in words2.
2. For each word in words1
   - Find the frequency of each letter in the word.
   - Check if the frequency of each letter in the word is greater than or equal to the maximum frequency of each letter in words2.
3. Return the words in words1 that satisfy the condition.

# Complexity Analysis
## Time Complexity O(n*m + k*j)
- n is the number of words in words1.
- m is the maximum length of the words in words1.
- k is the number of words in words2.
- j is the maximum length of the words in words2.
- For both iterations of words1 and words2, we need to iterate through the letters of the words, so the time complexity is O(n*m + k*j).

## Space Complexity O(1)
- Although we use multiple frequencies slice to store the frequency of each letter, the space complexity is O(1) because the size of the slice is constant (26).