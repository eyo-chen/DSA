# Problem Explanation

## Brute Force
This is the most straightforward approach. We iterate through the array and check if there exists a triplet (i, j, k) such that nums[i] < nums[j] < nums[k].

### Complexity Analysis
#### Time Complexity O(n^3)
#### Space Complexity O(1)

## Find Lower and Higher
This approach is a bit more efficient than the brute force solution, and it's still easy to understand.<br>

The idea is to pick each element as a potential middle element<br>
For each middle element, we:<br>
- Look left to find if there's any smaller number
- Look right to find if there's any larger number

If both left and right have numbers that are smaller and larger than the middle number, we have found a triplet.

### Complexity Analysis
#### Time Complexity O(n^2)
#### Space Complexity O(1)

## Greedy
This is the most efficient approach. It's also the most difficult to understand.<br>

Let me explain this elegant O(n) solution step by step.

**Core Idea:**
- `i` tracks the smallest number we've seen
- `j` tracks the second smallest number we've seen (but must come after a smaller number)
- If we find a number larger than both `i` and `j`, we've found our triplet

**Why It Works:**
1. We're essentially trying to maintain a sequence of two numbers (`i` and `j`), where `i < j`
2. If we find a number larger than both, it completes our triplet
3. Even if we update `i` after finding `j`, it doesn't invalidate our solution because the original `i` that led to our current `j` still exists in the array

**Let's walk through an example:**
Consider the array: `[2, 1, 5, 0, 4, 6]`

```
Initial state: i = MaxInt, j = MaxInt

Step 1: n = 2
i = 2, j = MaxInt    (2 becomes our smallest number)
 i
[2, 1, 5, 0, 4, 6]

Step 2: n = 1
i = 1, j = MaxInt    (found smaller number, update i)
    i
[2, 1, 5, 0, 4, 6]

Step 3: n = 5
i = 1, j = 5         (5 > i, so it becomes our j)
    i  j
[2, 1, 5, 0, 4, 6]

Step 4: n = 0
i = 0, j = 5         (found smaller number, update i)
       j  i
[2, 1, 5, 0, 4, 6]

Step 5: n = 4
i = 0, j = 4         (4 > i but < j, update j to smaller value)
          i  j
[2, 1, 5, 0, 4, 6]

Step 6: n = 6
Found triplet!       (6 > j > i)
          i  j
[2, 1, 5, 0, 4, 6]
```

**Why updating i doesn't break the solution:**
Let's look at what happens at step 4 when we update `i` from 1 to 0:
1. At this point, we already know that sequence `[1, 5]` exists in our array
2. When we update `i` to 0, we're essentially saying "we found an even better first number"
3. Even though we update `i`, the fact that we found a valid `j` (5) means there must have been a smaller number before it
4. So when we finally find 6, we know for sure that there exists some number < 4 < 6 in our sequence

Let's explain this crucial part in more detail using the same example `[2, 1, 5, 0, 4, 6]`.

When we update `i` to 0 (which comes after 5), it might seem like we're breaking the requirement that i < j < k in terms of indices. However, this still works because:

1. When we found `j = 5`, we know for certain that:
   - There exists some number (in this case, 1) that is less than 5
   - This smaller number appears before 5 in the array

2. So even though we later update `i = 0`:
   - The original sequence `[1, 5]` is still preserved in our array
   - Finding `i = 0` just gives us another potential starting point
   - But we don't need to use this new `i = 0` if we find our answer using the original sequence

Let's see what happens when we find 6:
````
Original sequence we found: [1, 5]
                           ^  ^
[2, 1, 5, 0, 4, 6]
      ^  ^     ^
      i  j     k

Even after updating to i = 0:
[2, 1, 5, 0, 4, 6]
      ^     ^  ^
      i     i  j

When we find 6, we have multiple valid triplets:
1. [1, 5, 6] - using our original sequence
2. [1, 4, 6] - using the updated j
3. [0, 4, 6] - using the updated i
````

The key insight is:
- We don't care about the exact position of `i`
- We only care that there exists some number smaller than `j`
- Once we've found a valid `j`, we know for sure there must be a smaller number before it
- So when we find a number larger than both `i` and `j`, we can guarantee a valid triplet exists

This is why the algorithm is considered "greedy" - it maintains just enough information (the existence of a smaller number) without needing to track the exact position or value of all potential candidates.

Think of it like this: if you've found a pair of increasing numbers anywhere in the array, finding a third larger number automatically creates a valid triplet, regardless of where your current `i` pointer is, because the original smaller number that led to your `j` still exists in the array.

**Key Insights:**
1. If we find a number that's greater than both `i` and `j`, we've definitely found a valid triplet because:
   - `j` only gets updated when we find a number greater than `i`
   - Therefore, somewhere in the array, there must be a sequence of three increasing numbers

2. The beauty of this solution is that we don't need to know exactly where our first number is - we just need to know that it exists somewhere before our current position.

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)
