# Problem Description
Given an array of string, return the distance between the nearest repeated strings. If no entry is repeated, return -1.

Input-Output:<br>
Example 1: 
Input:
```
[
  "This",
  "is",
  "a",
  "sentence",
  "with",
  "is",
  "repeated",
  "then",
  "repeated"
]
```
Output: 2<br>
Explanation: "repeated" (index 6) and "repeated" (index 8) are 2 positions away.<br>

Example 2<br>
Input:
```
[
  "This",
  "is",
  "a"
]
```
Output: -1<br>
Explanation: There are no repeated entries.<br>

# Problem Explanation
The problem is pretty simple. We need to find the nearest repeated entries in an array of strings. We can use a hash table to store the index of each string. When we encounter a string that is already in the hash table, we calculate the distance between the current index and the index of the string in the hash table. We keep track of the minimum distance. If no string is repeated, we return -1.

# Complexity Analysis
## Time Complexity O(n)
- We iterate through the array once to store the index of each string in the hash table.

## Space Complexity O(n)
- We use a hash table to store the index of each string.

