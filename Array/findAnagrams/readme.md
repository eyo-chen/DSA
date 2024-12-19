# Problem Explanation

## Use Two Hash Table And Generate Key
The core idea is very similar to the problem groupAnagrams.<br>
We first build a hash table for the pattern string p, and then generate a key for it.<br>
Then, we iterate through the string s, and maintain a current window of length len(p).<br>
For each window, we generate a key for it, and compare it with the key of the pattern string.<br>
If they are the same, we add the start index of the current window to the result.<br>

### Complexity Analysis
#### Time Complexity O(n + m)
- where n is the length of the string s, and m is the length of the pattern string p.
- we separately iterate through the string s and the pattern string p, and the time complexity is O(n + m).
- `genKey` function is O(1) because it only iterates through the hash table once.

#### Space Complexity O(1)
- We use two hash tables, and the space is only 26 * 2 = 52, which is constant.


## Use Sliding Window And Generate Key
The idea is very similar to the above solution, but we do not generate a key for the current window.<br>
Instead, we directly compare the hash table of the current window with the hash table of the pattern string.<br>
If they are the same, we add the start index of the current window to the result.<br>

Also, note that the way we maintain the window is a little bit different from the above solution.<br>

### Complexity Analysis
#### Time Complexity O(n + m)
- where n is the length of the string s, and m is the length of the pattern string p.
- we separately iterate through the string s and the pattern string p, and the time complexity is O(n + m).
- `arrayEqual` function is O(1) because it only iterates through the hash table once.

#### Space Complexity O(1)
- We use two hash tables, and the space is only 26 * 2 = 52, which is constant.


## Use Match Count
The idea is to use a match count to check if the current window is an anagram of the pattern string.<br>
The purpose of using a match count is to how many characters in the current window are correctly matched.<br>
When the match count is equal to the length of the pattern string, we add the start index of the current window to the result.<br>

How to update the match count?<br>
- When we add a character to the current window, we decrease the count of the character in the hash table of the pattern string.
- If the count is greater than 0, it means the character is in the pattern string, so we increase the match count.
- When we remove a character from the current window, we increase the count of the character in the hash table of the pattern string.
- If the count is greater than 0, it means the character is in the pattern string, so we decrease the match count.

### Complexity Analysis
#### Time Complexity O(n + m)
- where n is the length of the string s, and m is the length of the pattern string p.
- we separately iterate through the string s and the pattern string p, and the time complexity is O(n + m).

#### Space Complexity O(1)
- We use a hash table for the pattern string, and the space is only 26, which is constant.

