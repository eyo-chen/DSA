# Problem Explanation

## Brute Force Solution
The idea is to use recursion to solve the problem. We start from the top of the stairs and try to reach the bottom. We can either take 1 step or 2 steps at a time. So, we recursively try to reach the bottom from the current step.

We're first at the nth step, and ask "How many ways are there to get to the bottom from here if I can take 1 step or 2 steps at a time?"
We don't know the answer, so we ask the same question to the (n-1)th step and the (n-2)th step.
We keep asking this question recursively until we reach the bottom of the stairs.

Each recursive call represents a subproblem:<br>
How many ways are there to get to the bottom from here if I can take 1 step or 2 steps at a time?

Let's take an example to understand this better.<br>
n = 3<br>
- How many ways are there to get to the bottom from step 3? (step(3))
  - We don't know, so we ask the same question to step 2 and step 1.
- How many ways are there to get to the bottom from step 2? (step(2))
  - We don't know, so we ask the same question to step 1 and step 0.
- How many ways are there to get to the bottom from step 1? (step(1))
  - We don't know, so we ask the same question to step 0.
- How many ways are there to get to the bottom from step 0?
  - There is 1 way to get to the bottom from step 0, which is to stay there.

Now, we know that there is 1 way to get to the bottom from step 0, which is to stay there.<br>
So, we return 1 for the base case.<br>

Later, we get the answer for step(1) and step(2) from our recursive calls.<br>
step(1) = step(0) = 1<br>
step(2) = step(1) + step(0) = 1 + 1 = 2<br>
Now, we can calculate step(3) = step(2) + step(1) = 2 + 1 = 3<br>
So, there are 3 ways to get to the bottom from step 3.

Or, we can think of it as a tree:

```
                           step(3)
                   (1) /                    \ (2)
                step(2)                    step(1)
            (1)  /    \(2)             (1) /    \
            step(1)  step(0)          step(0)
            (1) /  
            step(0)
```
From this tree, we know there are 3 ways to get to the bottom from step 3.<br>
1 + 1 + 1 = 3<br>
1 + 2 = 3<br>
2 + 1 = 3<br>

Also note that we can count different combination as unique path.<br>
For example, 1 + 2 and 2 + 1 are different paths.

### Complexity Analysis
#### Time Complexity O(2^n)
- where n is the number of stairs.
- Each recursive call branches out to 2 more recursive calls
- And the depth of the tree is n
- So, the time complexity is O(2^n)

#### Space Complexity O(n)
- where n is the number of stairs.
- The depth of the recursive call stack is n
- So, the space complexity is O(n)


## Memoization Solution
From the above recursive tree, we can see that there are many overlapping(duplicate) subproblems.<br>
We count step(1) multiple times in our recursive calls.<br>
We can use a memoization hash map to store the result of each subproblem and reuse it instead of recalculating it.

In this solution, each subproblem is computed only once and then stored in the memo.<br>
After the first computation, subsequent calls for the same subproblem return the memoized result in O(1) time.

### Complexity Analysis
#### Time Complexity O(n)
- where n is the number of stairs.
- There are n+1 subproblems (0 to n).
- Each subproblem is solved exactly once and memoized.
- Each subproblem takes O(1) time to solve (excluding recursive calls).
- Therefore, the total time complexity is O(n+1) * O(1) = O(n).

#### Space Complexity O(n)
- where n is the number of stairs.
- The memo hash map stores the result of each subproblem.
- The maximum depth of the recursive call stack is n.
- So, the space complexity is O(n).

## Tabulation Solution
This is a bottom-up approach to solve the problem. We start from the bottom of the stairs and try to reach the top. We can either take 1 step or 2 steps at a time. So, we iteratively try to reach the top from the current step.

For each step, we ask two questions:
1. How many ways are there to get to the top from here if I can take 1 step?
2. How many ways are there to get to the top from here if I can take 2 steps?

Let's take an example to understand this better.<br>
n = 3<br>
First, create an array to store the number of ways to get to each step.<br>
[0, 0, 0, 0]<br>
For each index, it represents the number of ways to get to that step.

For the 0th step, there is 1 way to get to the top, which is to stay there.
So, table[0] = 1

- We're at the 1st step.
  - How many ways are there to get to the 1st step from the bottom?
  - I don't know, so I ask the same question to the 0th step.
  - The 0th step says there is 1 way to get to the top from here, which is to stay there.
  - So, table[1] = table[0] = 1

- We're at the 2nd step.
  - How many ways are there to get to the 2nd step from the bottom?
  - I don't know, so I ask the same question to the 1st step and the 0th step because I can take 1 step or 2 steps at a time.
  - table[2 - 1] = table[1] = 1
  - table[2 - 2] = table[0] = 1
  - So, table[2] = table[1] + table[0] = 1 + 1 = 2

- We're at the 3rd step.
  - How many ways are there to get to the 3rd step from the bottom?
  - I don't know, so I ask the same question to the 2nd step and the 1st step because I can take 1 step or 2 steps at a time.
  - table[3 - 1] = table[2] = 2
  - table[3 - 2] = table[1] = 1
  - So, table[3] = table[2] + table[1] = 2 + 1 = 3

For each step, we sum up the result from the following questions:
1. How many ways are there to get to the top from here if I can take 1 step?
2. How many ways are there to get to the top from here if I can take 2 steps?
  

### Complexity Analysis
#### Time Complexity O(n)
- where n is the number of stairs.
- We iterate through each step from 0 to n.
- For each step, we ask two questions and sum up the result.
- So, the time complexity is O(n).

#### Space Complexity O(n)
- where n is the number of stairs.
- We use an array to store the number of ways to get to each step.
- The size of the array is n+1.
- So, the space complexity is O(n).