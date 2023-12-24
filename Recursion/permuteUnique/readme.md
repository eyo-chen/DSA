# Problem Explanation

This problem is very similar to permutation 1<br/>
but the main difference is how do we handle duplicate element in the input array

Let's try to use the solution of permutation 1 in this problem, and see the outcome<br/>
For example, nums = [1,1,2], for the sake of differentiation, we use [1a, 1b, 2] to represent
<pre>
                                               [1a, 1b, 2]
                        1a                          1b                    2
            1b                  2         1a                2      1a          1b
            2                   1b        2                1a      1b          1a
</pre>
The result would be like this<br/>
[1a, 1b, 2], [1a, 2, 1b], [1b, 1a, 2], [1b, 2, 1a], [2, 1a, 1b], [2, 1b, 1a]<br/>
[1,  1,  2], [1,  2,  1], [1,  1,  2], [1,  2,  1], [2,  1,  1], [2,  1,  1]

Look at the recurisve tree<br/>
Where start the duplication?<br/>
It's the start point when we choose 1b<br/>
=> 1b -> 1a -> 2 (second level)<br/>
=> 1b -> 2 -> 1a (second level)<br/>
=> 2 -> 1b -> 1a (second level)<br/>

Look closly, ***duplication happens when we choose the second duplicate element without choosing the first one***<br/>

So the core logic of this solution is<br/>
***When we see the second duplicate element, and the first duplicate element is not chosen yet, we skip it***<br/>

To correctly implement this logic, we need to ***sort the input array first***<br/>

## Choices and Constraints

- **Choice:** Choose element in the decision space
- **Constraint:** Have to which element we can choose and which element we can't choose + When we see the second duplicate element, and the first duplicate element is not chosen yet, we skip it
- **Goal:** The length of working permutation is equal to the length of input

## Caveats
### Caveat 1: Initialize vector before accessing it
```c++
if (i > 0 && nums[i] == nums[i - 1] && !used[i - 1]) continue;
```
- This line represents the logic of skipping the second duplicate element
- `i > 0` is to make sure we don't access `nums[-1]`

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


# Problem Explanation2
The second solution is more intuitive<br/>

***Convert input array to map, and use the map to keep track of the number of each element***<br/>
key -> each unique element in the input array<br/>
value -> the frequence of each element in the input array

For example, nums = [1,1,2]<br/>
<pre>
{
 1: 2
 2: 1
}
</pre>

Instead of looping input array, we loop through the map<br/>
The main benefit is that we won't have duplicate element in the map<br/>
So we don't need to worry about the duplicate element<br/>

## Recursive Tree Visualization
<pre>
                                       {1: 2, 2: 1}
                                1                         2
                            {1:1, 2:1}                {1:2,2:0}
                  1                      2                  1
              {1:0,2:1}             {1:1,2:0}           {1:1,2:0}
                2                      1                   1
              {1:0,2:0}             {1:0,2:0}            {1:0,2:0}
</pre>

# Complexity Analysis

n = the legnth of input string

## Time Complexity: O(n * n!)
- The time complexity is the same as the first solution
- Because the worst case is each element in the input array is unique

## Space Complexity: O(n)
- The space complexity is the same as the first solution