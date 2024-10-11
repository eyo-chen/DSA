# Problem Description

Write a function that takes in a list of unique, lowercase strings and returns a list of semordnilap pairs. A semordnilap pair is defined as a set of two different strings where the reverse of one word is the same as the forward version of the other. The order of the returned pairs and the order of the strings within each pair does not matter.

Input:<br>
words = ["diaper", "abc", "test", "cba", "repaid"]<br>
Output: [["diaper", "repaid"], ["abc", "cba"]]<br>
Explanation:<br>
"diaper" and "repaid" spell each other backwards<br>
"abc" and "cba" spell each other backwards<br>

# Problem Explanation
This problem is straightforward.

We can use a hash table to store the strings we have seen.<br>
Then we can loop through the list of strings, reverse each string, and check if the reversed string is in the hash table.<br>
If it is, and it's not the same as the original string, we have found a semordnilap pair.

Note that We use a hash table (implemented as a map in Go) to efficiently check if we've seen a string before, and to avoid adding duplicate pairs to the result.


# Complexity Analysis
## Time Complexity O(N * M)
- N is the number of strings in the input list
- M is the average length of the strings
- We loop through the list of strings once, and for each string, we reverse it and check if the reversed string is in the hash table
- So the time complexity is O(N * M)

## Space Complexity O(N)
- We use a hash table to store the strings we have seen, so the space complexity is O(N)