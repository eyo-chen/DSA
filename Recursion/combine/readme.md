# Problem Explanation

## Choices and Constraints

- **Choice:** Choose element from i ~ n
- **Constraint:** After choosing element i, we can only choose element from i + 1 ~ n
- **Goal:** The length of working combination is equal to k

## Recursive Tree Visualization
Think the recursive tree like this when n = 4, k = 2
<pre>
                                          1,2,3,4
                    1              2                3             4
         2      3       4       3    4             4  
</pre>
After choosing 1, we can only choose 2, 3, 4. So we have 3 choices.<br/>
After choosing 2, we can only choose 3, 4. So we have 2 choices.<br/>
After choosing 3, we can only choose 4. So we have 1 choice.<br/>

When the depth of recursive tree is equal to k, we know we're done.

# Complexity Analysis

n = the input n<br/>
k = the input k

## Time Complexity: O((n ^ k) * k)
- The largest branching factor is n
- The deepest height of recursive tree is k
- When pushing tmp to ans, it takes O(k) time because we have to copy the whole vector to the answer vector
- This is just the roughly tight bound because we know the branching factor is not exactly n for each caes

## Space Complexity: O(k)
- The deepest height of recursive tree is k
