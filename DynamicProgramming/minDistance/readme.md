# Recursive Approach
At first glance, this problem seems extremely difficult. However, if we start by solving it with small and simple inputs and try to identify patterns, this problem becomes much easier to understand.

## Understand Base Cases

Before diving into the complex cases, let's understand when our recursion should stop. These base cases are crucial for the recursive solution.

- **Base Case 1:** When we've processed all characters in word1 (i == len(word1))  
Let's say we're converting "cat" to "catch". We compare character by character:
  - Position 0: 'c' matches 'c' ✓
  - Position 1: 'a' matches 'a' ✓  
  - Position 2: 't' matches 't' ✓

Now we've reached the end of "cat" (i = 3), but "catch" still has "ch" remaining (j = 3, but len("catch") = 5).  
What do we need to do? We need to **insert** all remaining characters "ch" from "catch".  
That's exactly `len(word2) - j = 5 - 3 = 2` operations.

- **Base Case 2:** When we've processed all characters in word2 (j == len(word2))  
Now let's consider converting "catch" to "cat". We compare:
  - Position 0: 'c' matches 'c' ✓
  - Position 1: 'a' matches 'a' ✓
  - Position 2: 't' matches 't' ✓

At this point, we've used up all characters in "cat" (j = 3), but "catch" still has "ch" remaining (i = 3).  
What do we need to do? We need to **delete** all remaining characters "ch" from "catch".  
That's exactly `len(word1) - i = 5 - 3 = 2` operations.

**Key Insight:** These base cases handle the "cleanup" phase when one word runs out before the other. Think of it like copying a recipe - if one ingredient list is longer than the other, you either need to add the missing ingredients or remove the extra ones.

## Finding the Pattern

- **Example 1:** w1: "abc", w2: ""  
What's the minimum number of operations needed to convert "abc" to ""?  
Simple! Just delete all characters from "abc". So the answer is 3.

- **Example 2:** w1: "", w2: "abc"  
What's the minimum number of operations needed to convert "" to "abc"?  
Simple! Just insert all characters from "abc". So the answer is 3.

- **Example 3:** w1: "abf", w2: "cde"  
What's the minimum number of operations needed to convert "abf" to "cde"?  
This isn't as straightforward as the previous examples, so let's break it down step by step.  
First, we place two pointers at the beginning of the strings "abf" and "cde".

```
i
a  b  f

k
a  d  e
```

i = 0, k = 0  
Both characters are the same, so we can simply ignore the current characters and ask the subproblem: "What's the minimum number of operations needed to convert "bf" to "de"?"  
In other words, we move both pointers forward.

```
   i
a  b  f

   k
a  d  e
```

i = 1, k = 1  
Now, the two characters are different: 'b' and 'd'.  
We have three operations to try:
1. Replace
2. Insert  
3. Delete

Let's examine them one by one.

### Replace
When we want to replace, we obviously need to replace the character with the target character.  
In this example, we need to replace 'b' with 'd': "abf" → "adf"  
After replacement, how do we update the two pointers?  
***We move both pointers forward***  
This means "We've performed one operation (Replace), so now let's consider the rest of the string."

```
      i
a  d  f

      k
a  d  e
```


### Insert
When we want to insert, we obviously need to insert the target character before the current character.  
In this example, we need to insert 'd' before 'b': "abf" → "adbf"  
Does this mean we actually need to insert a character into our string?  
Not necessarily! Let's look at the two strings side by side:  
"adbf"  
"ade"  

We can see that the "ad" part is exactly the same after insertion.  
So, we can essentially ask the subproblem: "What's the minimum number of operations needed to convert "bf" to "e"?"

Recall the position of the two pointers:

```
   i
a  b  f

   k
a  d  e
```

i = 1, k = 1

After insertion:

```
      i
a  d  b  f

      k
a  d  e
```

If we ignore the newly inserted character "d", pointer `i` essentially remains the same.  
***We only need to update pointer `k`***  
This means "Let's consider the rest of strings `bf` and `e` after insertion."

### Delete
When we want to delete, we obviously need to delete the current character.  
In this example, we need to delete 'b': "abf" → "af"  
Similar to insertion, we don't actually need to delete the current character.  
We can just "ignore" it and consider the rest of the string.  
After deletion, how do we update the two pointers?  
***We only need to update pointer `i`***  
This means "Let's consider the rest of strings `f` and `de` after deletion."

### Summary
Let's summarize how to update the two pointers for each operation:
1. **Replace:** Move both pointers forward (i++, k++)
2. **Insert:** Update pointer `k` (i, k++)  
3. **Delete:** Update pointer `i` (i++, k)

One important thing to remember: for each operation, we need to add 1 to the total operation count.

## Complexity Analysis

### Time Complexity O(3^(m+n))
- where m and n are the lengths of word1 and word2 respectively.  
- In the worst case, at each recursive call, we explore all three possible operations (insert, delete, replace). The maximum depth of recursion is m + n (when we need to process all characters from both strings), leading to exponential time complexity.

### Space Complexity O(m+n)
- The maximum depth of the recursion stack is m + n in the worst case, where m and n are the lengths of the two input strings.

# DP Approach

Once we understand the pattern from the recursive approach, implementing a dynamic programming solution becomes straightforward. We essentially build a standard DP table and use the pattern (formula) we discovered to fill it in.

First, let's build the DP table:

```
     ""   h    o   r   s   e
""   0    0    0   0   0   0
r    0    0    0   0   0   0
o    0    0    0   0   0   0
s    0    0    0   0   0   0
```

Now, how do we initialize the DP table?

**Consider the first row:**  
Each cell represents: "What's the minimum number of operations needed to convert "", "h", "ho", "hor", "hors", "horse" to an empty string ""?"  
The answer is simply the column index of that cell.  
For example, we need exactly two operations (Delete) to convert "ho" to "". So we set dp[0][2] = 2.  
For example, we need exactly three operations (Delete) to convert "hor" to "". So we set dp[0][3] = 3.

```
     ""   h    o   r   s   e
""   0    1    2   3   4   5
r    0    0    0   0   0   0
o    0    0    0   0   0   0
s    0    0    0   0   0   0
```

**Consider the first column:**  
Each cell represents: "What's the minimum number of operations needed to convert "" to "", "r", "ro", "ros"?"  
The answer is simply the row index (i) of that cell.  
For example, we need exactly one operation (Insert) to convert "" to "r". So we set dp[1][0] = 1.  
For example, we need exactly two operations (Insert) to convert "" to "ro". So we set dp[2][0] = 2.  
For example, we need exactly three operations (Insert) to convert "" to "ros". So we set dp[3][0] = 3.

```
     ""   h    o   r   s   e
""   0    1    2   3   4   5
r    1    0    0   0   0   0
o    2    0    0   0   0   0
s    3    0    0   0   0   0
```

Then, we use the following formula to fill each cell:

- **When two characters are the same** (`word1[r-1] == word2[c-1]`), we simply copy the diagonal value: `dp[r][c] = dp[r-1][c-1]`
  - This means "The current characters are the same, so we can ignore them and see what's the minimum number of operations needed for the remaining substrings."

- **When two characters are different** (`word1[r-1] != word2[c-1]`), we need to consider all possible operations:
  - **Replace:** `dp[r][c] = dp[r-1][c-1] + 1`
  - **Insert:** `dp[r][c] = dp[r][c-1] + 1`  
  - **Delete:** `dp[r][c] = dp[r-1][c] + 1`
  - We take the minimum value among all possible operations.

## Complexity Analysis

### Time Complexity O(m × n)
- where m and n are the lengths of word1 and word2 respectively.  
- We fill a 2D DP table of size (m+1) × (n+1), and each cell takes constant time to compute.

### Space Complexity O(m × n)
- We use a 2D DP table to store the minimum edit distances for all substrings of word1 and word2. The space complexity is therefore O(m × n).