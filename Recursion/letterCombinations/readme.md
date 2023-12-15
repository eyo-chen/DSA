# Problem Explanation

Illustrate this problem as "decision space"

For example, input digits is "23"
We can think something like following

2 3<br/> 
_ _ -> decision space <br/>
a d -> <br/>
b e      -> these three are all possible decision space<br/>
c f -> <br/>

For each slot or decision space, we have three choices
- a, b, c for the first slot
- d, e, f for the second slot


## Choices and Constraints

- **Choice:** The characters mapping to the digits
- **Constraint:** None
- **Goal:** Find all possible combinations

## Recursive Tree Visualization
Start at first index, "2", we know it can have "abc"
<pre>
                               "23"
                "a"            "b"            "c"
</pre>
First index "2" can have three decisions<br/> 
Then, each decisions can have another more decisions
<pre>
                               "23"
                "a"             "b"            "c"
     "ab"    "ad"   "ae"  "bd"   "ba"  "be"  "cd"  "ca"  "ce"
</pre>

# Complexity Analysis

n = the legnth of input digits

## Time Complexity: O(4^n)
- Branching Factor = 4
   - For each single digit, we'll explore four decisions at most
- Depth = n
    - If the length of input is 4, then we'll have 4 height of the tree
- Each call stack = O(1)
    - Because we only do constant work

## Space Complexity: O(n)