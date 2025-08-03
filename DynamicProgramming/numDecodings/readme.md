# Problem Explanation

## Top-Down Solution

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

### Complexity Analysis
#### Time Complexity O(n)
- where n is the length of the string.
- because we use memoization, we don't calculate the result of the subproblem again.
- each subproblem will be calculated once.
- the time complexity would be O(n^2) if we don't use memoization.
- because the branching factor is 2, and the depth of the tree is n.

#### Space Complexity O(n)
- because we use memoization, we need to store the result of the subproblem.
- the space complexity would be O(n) to store the memoization table.


## Bottom-Up Solution

In the top-down approach, we started with the full string "123" and asked "How many ways can we decode this?" Then we broke it down into smaller subproblems.<br>
The bottom-up approach reverses this thinking entirely. Instead of starting with the big problem and breaking it down, we start with the smallest possible problems and build our way up to the final answer.

Think of it this way: instead of asking "How many ways can we decode from position i to the end?"<br>
We now ask ***"How many ways can we decode from the beginning up to position i?"***<br>
This shift in perspective is the key to understanding bottom-up DP.

Let's work through the same example "123" but this time building from the ground up.<br>
We'll create a table where ***`dp[i]` represents "the number of ways to decode the first i characters of the string."***

Here's how we set up our foundation:
- `dp[0] = 1` represents decoding zero characters (empty string), which has exactly one way: do nothing
- `dp[1]` will represent the number of ways to decode just the first character
- `dp[2]` will represent the number of ways to decode the first two characters
- `dp[3]` will represent the number of ways to decode all three characters (our final answer)


Let's trace through building our DP table for "123":

**Step 1: Initialize the base case**
```
String: "123"
Index:   0 1 2 3
dp:     [1, ?, ?, ?]
```
We set `dp[0] = 1` because there's one way to decode an empty string.

**Step 2: Handle the first character "1"**
```
String: "123"
Index:   0 1 2 3
dp:     [1, 1, ?, ?]
```
Since "1" is a valid single character (maps to 'A'), we set `dp[1] = 1`. There's exactly one way to decode the string "1".

**Step 3: Handle the first two characters "12"**
```
String: "123"
Index:   0 1 2 3
dp:     [1, 1, 2, ?]
```
Now we're asking: "How many ways can we decode '12'?"<br>
We can arrive at this position in two ways:<br>
- From `dp[1]`: Take the 1 way to decode "1", then add "2" as a single character → "A" + "B" = "AB"
- From `dp[0]`: Take the 1 way to decode nothing, then add "12" as a two-character code → "L"

So `dp[2] = dp[1] + dp[0] = 1 + 1 = 2`. There are 2 ways to decode "12": "AB" and "L".

**Step 4: Handle all three characters "123"**
```
String: "123"
Index:   0 1 2 3
dp:     [1, 1, 2, 3]
```
Finally, we ask: "How many ways can we decode '123'?" Again, we can arrive here in two ways:
- From `dp[2]`: Take the 2 ways to decode "12" ("AB" and "L"), then add "3" → "ABC" and "LC"
- From `dp[1]`: Take the 1 way to decode "1" ("A"), then add "23" as a two-character code → "AW"

So `dp[3] = dp[2] + dp[1] = 2 + 1 = 3`. There are 3 ways to decode "123": "ABC", "LC", and "AW".


The pattern we see is that at each position i, we're asking: "How could I have arrived here?" There are exactly two possibilities:

1. **Single character decode**: I came from position i-1 by decoding one character
   - This is valid if the character at position i-1 is between '1' and '9'
   - If valid, I inherit all the ways from `dp[i-1]`

2. **Two character decode**: I came from position i-2 by decoding two characters together
   - This is valid if the two-character substring forms a number between 10 and 26
   - If valid, I inherit all the ways from `dp[i-2]`

This gives us our recurrence relation: `dp[i] = dp[i-1] + dp[i-2]` (with appropriate validity checks).

### Complexity Analysis
#### Time Complexity O(n)
- where n is the length of the string
- we iterate through the string once, and for each position, we perform constant time operations
- each position in our DP table is calculated exactly once

#### Space Complexity O(n)
- we need to store the DP table with n+1 entries
- this can be optimized to O(1) since we only need to look back at most 2 positions
- the space-optimized version uses only a few variables instead of the entire table