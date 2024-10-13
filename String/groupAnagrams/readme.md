# Problem Explanation

## Brute Force
The idea to use brute force is to iterate through the array and for each word, check if it is an anagram of any other word in the array. If it is, add it to the current group.<br>
For example, if we have the following words: ["eat", "tea", "tan", "ate", "nat", "bat"]<br>
We take "eat" and loop through the rest of the array, for each word, we ask if it is an anagram of "eat". If it is, we add it to the current group.<br>
After we are done with "eat" and it's anagrams, we need to add those to seen hash table<br>
So that when we see those words has been added to a group, we skip it.

How can we check if two words are anagrams?<br>
We just need to generate a frequency count array for each word, and compare those two frequency count arrays.<br>
If any element in the frequency count array is different, then the two words are not anagrams.

### Complexity Analysis
#### Time Complexity O(n^2 * m)
- n is the number of words in the array
- m is the maximum length of the words
- We have nested loop, the outer loop run n times, and the inner loop also run n times in the worst case.
- For each nested loop, we have a function call to generate the two frequency count array, which is O(m)

#### Space Complexity O(n)
- We use a hash table to store the words that have been seen, the space is O(n)
- The space for the frequency count array is O(1) because it's a constant size of 26.


## Optimized
The idea is to use a hash table to group the words by their frequency count key.<br>
For example, if we have a word "eat", the frequency count array is [1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0], the key is "1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0".<br>
We convert the frequency count array to a string as the key, so that when we encounter the same frequency count array, we can find it in the hash table.

For each word,<br>
1. Generate the frequency count key
2. Add the word to the group that is associated with the key
3. Return all the groups in the hash table

Note that there are two pitfalls in this problem:<br>
We might try to use the sum of int value of the frequency count array as the key<br>
This is wrong because 2 + 2 + 2 = 3 + 3, so it's hard to distinguish if two words are anagrams.<br>
Two words might have same sum of frequency count array, but they are not anagrams.<br>
That's why we can't use the sum of frequency count array as the key.

The other pitfall is that we might try to not use "," to join the frequency count array as the key<br>
This is also wrong because 1000 might means "1,0,0,0" or "10,0,0"<br>
We need to use "," to differentiate the difference between "1,0,0,0" and "10,0,0"
That's why we need to use "," to join the frequency count array as the key.

### Complexity Analysis
#### Time Complexity O(n * m)
- n is the number of words in the array
- m is the maximum length of the words
- We have a loop that iterates through the array once, and for each word, we generate the frequency count key and add the word to the group, all of which are O(m) operations.

#### Space Complexity O(n * m)
- We use a hash table to store the groups, the space is O(n)