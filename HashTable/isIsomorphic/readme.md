# Problem Explanation

The key to solve this problem is to note that<br>
***Each character must map to exactly one other character, and no two characters may map to the same character.***<br>
***The mapping must be bi-directional, meaning if a <-> b, then b <-> a.***<br>
For example, s = "bad" and t = "bab" is not isomorphic<br>
i = 0, we map b <-> b<br>
i = 1, we map a <-> a<br>
i = 2, we map d <-> b, but b is already mapped to a, so it's not isomorphic<br>

Therfore, in order to map be-directional, we need to use two hash tables, one for s -> t, and one for t -> s.<br>
Iterate through the string, and for each character, check if the character is already mapped in the hash table.<br>
If it is, check if the mapped character is the same as the character in the other string.<br>
Return false if it is not<br>
Otherwise, add the mapping to the hash table.<br>
If we finish iterating through the string, return true.<br>

# Complexity Analysis
## Time Complexity O(n)
## Space Complexity O(n)
