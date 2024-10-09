# Problem Explanation

## Hash Table
The main idea is to count the number of each character in the string.<br>
Then, we can use the count to calculate the longest palindrome.

When we loop through the hash table, we have to consider the following:
- If the count of the character is odd, we need to set the `hasOdd` to `true`.
  - In the end, we need to check if `hasOdd` is `true`. If so, we need to add 1 to the result.
  - For example, if the string is "abccccdd", the hash table will be `{a: 1, b: 1, c: 4, d: 2}`.
  - After we adding 4 + 2 (c + d), we have 6.
  - We need to add 1 to the result to make it 7. It could either be `a` or `b`.
- If the count of the character is greater than or equal to 2, we can add `count / 2 * 2` to the result.
  - This is because we can use the character to form a palindrome.
  - Note that we both consider the even and odd count of the character.
  - For even count, we just add all the count to the result.(4/2 * 2 = 4)
  - For odd count, we add the largest even number to the result.(5/2 * 2 = 4)

### Complexity Analysis
#### Time Complexity O(n)
- where n is the length of the string.

#### Space Complexity O(1)
- Because we deal with the character, and there are only 52 characters(uppercase and lowercase).


## HashSet
The main idea is to use a hash set to store the character.<br>
It's similar to the hash table, but the code is simpler, and we only have one pass.<br>

When looping through the string, we have to consider the following:
- If the character is already in the hash set
  - If yes, we add 2 to the result, and remove the character from the hash set.
    - We add 2 to the result because we find a PAIR to form a palindrome.
    - After we find a pair, we need to remove the character from the hash set.
  - If no, we add the character to the hash set.


In the end, if the hash set is not empty, we add 1 to the result.<br>
After iteration, if there is any character in the hash set, it means there is a odd count character in the string.<br>
We just need to add 1 to the result.

### Complexity Analysis
#### Time Complexity O(n)
- where n is the length of the string.

#### Space Complexity O(1)
- Because we deal with the character, and there are only 52 characters(uppercase and lowercase).
