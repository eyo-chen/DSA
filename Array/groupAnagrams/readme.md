# Problem Explanation

## Brute Force
The idea is to use a brute force approach to solve this problem.<br>
Suppose the input array is `["eat", "tea", "tan", "ate", "nat", "bat"]`<br>
We start from the first string `eat`<br>
Then we check the rest of the strings to see if they are anagrams of `eat`<br>
If they are, we add them to the current group<br>
Otherwise, we just skip them<br>
After checking the rest of the strings, we get a group `["eat", "tea", "ate"]`<br>
Now, how do we skip the strings that we already checked?<br>
We can use a boolean array to keep track of the strings that we already checked<br>
If a string is already checked, we just skip it<br>

Let's summarize the steps:
1. Initialize a boolean array to keep track of the strings that we already checked
2. Initialize an empty array to store the result
3. Iterate through the input array
4. For each string, check if it is already checked
5. If it is, skip it
6. If it is not, add it to the current group
7. Then, we check the rest of the strings to see if they are anagrams of the current string
8. If they are, we add them to the current group
9. Add the current group to the result
10. Return the result

### Complexity Analysis
#### Time Complexity O(n^2 * m)
- n is the number of strings in the input array
- m is the average length of the strings
- We have to check each string with the rest of the strings, so it's O(n^2)
- For each check, we need to compare the characters of the two strings, so it's O(m)

#### Space Complexity O(n)
- n is the number of strings in the input array
- We need to store the boolean array to keep track of the strings that we already checked, so it's O(n)

## Sorting
The idea is to sort each string and use the sorted string as the key in the hash table<br>
For example, the string `eat` will be sorted to `aet`, so we can use `aet` as the key to group the anagrams<br>
After grouping each anagram, we can get the result<br>

### Complexity Analysis
#### Time Complexity O(n * m log m)
- n is the number of strings in the input array
- m is the average length of the strings
- We need to sort each string, so it's O(m log m)
- We have to do this for each string, so it's O(n * m log m)

#### Space Complexity O(n)
- n is the number of strings in the input array
- We need to store the hash table to store the groups of anagrams, so it's O(n)

### Use Frequency As Key
The idea is to convert a string to a frequency array, and use the frequency array as the key in the hash table<br>
However, we can't directly use the frequency array as the key because the array is not a valid type in Go<br>
So we need to convert the frequency array to a string<br>
For example, the frequency array `[1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0]` will be converted to `"1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0"`<br>
Then we can use this string as the key to group the anagrams<br>
After grouping each anagram, we can get the result<br>

### Complexity Analysis
#### Time Complexity O(n * m)
- n is the number of strings in the input array
- m is the average length of the strings
- We need to convert each string to a frequency array, so it's O(m)
- We have to do this for each string, so it's O(n * m)

#### Space Complexity O(n)
- n is the number of strings in the input array
- We need to store the hash table to store the groups of anagrams, so it's O(n)
