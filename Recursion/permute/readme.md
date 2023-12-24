# Problem Explanation

In order to find all the permutation, have to really think about "decision space"

For example, nums = [1,2,3]<br/> 
We know that we're gonna have three slots to fill up <br/> 
__ __ __<br/>
=> At this point, our decision space is [1,2,3]<br/>
=> We have three choices at this state of call stack<br/>
=> Choose 1, 2 or 3<br/>

Use 1 to fill up the first slot<br/>
1 __ __<br/>
=> Once choosing the 1, now the decision space is [2,3]<br/>
=> Now we can only have two choices, we're not allowed to choose 1 anymore<br/>

Use 2 to fill up the second slot<br/>
1 2 __<br/>
=> Now the decision space is just [3]<br/>

1 2 3<br/>
=> Base case

Use 2 to fill up the first slot(after backtracking) <br/>
2 __ __<br/>
=> Once choosing the 2, now the decision space is [1,3]

so on and so forth....

The main difficulty is how to keep the decision space<br/>
Which means for any single state of call stack<br/>
we have to know what's our decision space<br/>
what element we can choose right now, and what element we can't choose

In this solution, i use array to keep track<br/>
For example, array = [false, false, false], it means decision space is [1,2,3]<br/>
__ __ __<br/>
=> At this point, array = [false, false, false]<br/>
=> decision space is [1,2,3]<br/>

Use 1 to fill up the first slot<br/>
1 __ __<br/>
=> Once choosing the 1, now the array = [true, false, false], which means<br/>
=> decision space is [2,3]<br/>
=> now we can only have two choices, we're not allowed to choose 1 anymore<br/>

Use 2 to fill up the second slot<br/>
1 2 __<br/>
=> array = [true, true, false]<br/>
=> Now the decision space is just [3]<br/>

1 2 3<br/>
=> Base case<br/>
=> array = [true, true, true]<br/>

Use 2 to fill up the first slot<br/>
2 __ __<br/>
=> array = [false, true, false]<br/>
=> Once choosing the 2, now the decision space is [1,3]<br/>

so on and so forth...

## Choices and Constraints

- **Choice:** Choose element in the decision space
- **Constraint:** Have to which element we can choose and which element we can't choose
- **Goal:** The length of working permutation is equal to the length of input

## Recursive Tree Visualization
Think the process of recursion in prermutation like a tree
<pre>
                                         [a, b, c]
                [a, _, _]                [b, _, _]                 [c, _, _]
        [a, b, _]      [a, c, _]   [b, a, _]     [b, c, _]    [c, a, _]   [c, b, _]
        [a, b, c]      [a, c, b]   [b, a, c]     [b, c, a]    [c, a, b]   [c, b, a]
</pre>

And the decision space(DS)(only left hand side)
<pre>
                                         [a, b, c]
                [a, _, _]         DS = [T,F,F]       
        [a, b, _]   DS = [T,T,F]    [a, c, _]    DS = [T,F,T]
        [a, b, c]   DS = [T,T,T]    [a, c, b]    DS = [T,T,T]
</pre>
=> As we could see, when [a, b, c] backtrack to [a, b, _], decision space also need to set back to false<br/>
=> Same thing for when [a, b, _] backtrack to [a, _, _]

## Caveats
### Caveat 1: Initialize vector before accessing it
```c++
vector<bool> used;
used[0] = true;
```
- This will cause runtime error(aka segmentation fault)
- It's important to initialize the vector before accessing it

# Complexity Analysis

n = the legnth of input string

## Time Complexity: O(n * n!)
- Look at the recursive tree visualization
- For first level, we have n choices
- For second level, we have n - 1 choices
- For third level, we have n - 2 choices
- So the branching factor is n * (n - 1) * (n - 2) * ... * 1 = n!
- `ans.push_back(tmp)` takes O(n) time because we have to copy the whole vector to the answer vector

## Space Complexity: O(n)
- The deepest height of recursive tree is n
