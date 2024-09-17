# Problem Explanation

Let's see how to solve this problem.<br>

We are given a string s "123", how many ways can we decode this string?<br>
I don't know the answer, but I do know we have two choices:<br>
1. Decode first character -> "1"
2. Decode first two characters -> "12"

After decoding first character "1", we are left with "23"<br>
Then we keep asking the same question, how many ways can we decode "23"?<br>
We have two choices again:
1. Decode first character -> "2"
2. Decode first two characters -> "23"

After decoding first character "2", we are left with "3"<br>
Then we keep asking the same question, how many ways can we decode "3"?<br>
We have only one choice:
1. Decode first character -> "3"

After decoding "3", we have an empty string, and there is only one way to decode an empty string -> "".<br>
We get the answer, and can return to the previous call with the answer "1".<br>

For each recursive call, it represents a subproblem.<br>
How many ways can we decode the current string?<br>
We only have two choices, decode one character or decode two characters.<br>
Or we know there's a base case, if we have an empty string, there is only one way to decode it -> "".<br>

Let's see the recursive tree:
```
                       "123"
                (1) /          \ (12)
               "23"             "3"
         (2)/     \(23)    (3)/
          "3"     ""         ""
      (3)/
        ""
```
In the implementation, we don't necessarily need to slice the string, we can just use an index to keep track of the current position in the string.

Now, let's see what's the base case:<br>
- If the index pointer is eqaul or greater than the length of the string, we return 1 because there's only one way to decode an empty string.
- If the current character is "0", we return 0 because "0" doesn't have any mapping.
  - This is the constraint of the problem.

Let's see the recursive case:<br>
- Decode one character:
  - We add the result of helper(s, index + 1) to the ways.
- Decode two characters:
  - We add the result of helper(s, index + 2) to the ways.
  - But we need to check if the two characters can be decoded.
  - Two characters can't be decoded if
    1. The index + 2 is greater than the length of the string.
       - For example, if the string is "1234", when the index is 3, we can't decode `"1234"[3:5]`, that's out of bound.
    2. The two characters are greater than 26.


Note that we might have overlapping subproblems.<br>
Let's see the recursive tree again:
```
                       "123"
                (1) /          \ (12)
               "23"             "3"
         (2)/     \(23)    (3)/
          "3"     ""         ""
      (3)/
        ""
```
When the index is at `2`, which is "3", we calculate the result twice.<br>
To optimize the solution, we can use memoization.<br>
We use index as the key to store the result of the subproblem.

# Complexity Analysis
## Time Complexity O(n)
- where n is the length of the string.
- because we use memoization, we don't calculate the result of the subproblem again.
- each subproblem will be calculated once.
- the time complexity would be O(n^2) if we don't use memoization.
- because the branching factor is 2, and the depth of the tree is n.

## Space Complexity O(n)
- because we use memoization, we need to store the result of the subproblem.
- the space complexity would be O(n) to store the memoization table.