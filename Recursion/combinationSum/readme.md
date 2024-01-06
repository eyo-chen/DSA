# Problem Explanation

Note that the regular recursion approach will cause the duplicate combination<br/>

For the sake of simplicity, let's assume that the input array is [2, 3] and the target is 7<br/>
<pre>

                                                                  7
                             [2]                                                                   [3]  
           [2,2]                     [2,3]                                     [3,3]                        [3,2]
    [2,2,2]     [2,2,3]        [2,3,3]    [2,3,2]                      [3,3,3]    [3,3,2]           [3,2,2]    [3,2,3]
</pre>
As we can see, there are duplicate combinations<br/>
For example, [2, 2, 3] and [2, 3, 2] are the same combination<br/>

So the key point of this problem is ***How to avoid the duplicate combination?***<br/>

The key idea is ***Once we choose an element, we can't choose the element before it***<br/>
For example, if we're given candidates [2, 3, 6, 7]<br/>
- If we choose 2, we can only choose 2, 3, 6, 7
- If we choose 3, we can only choose 3, 6, 7
- If we choose 6, we can only choose 6, 7
- If we choose 7, we can only choose 7

Why this idea works?<br/>
Look at the above recursive tree, when the duplicate combination happens?<br/>
When we already choose 3([2, 3]), and we're at the second level<br/>
We still go back to choose 2 again, and we get [2, 3, 2]<br/>
Which is the same combination as [2, 2, 3]<br/>

## Recursive Tree Visualization
<pre>

                                                                  7
                             [2]                                                                   [3]  
           [2,2]                     [2,3]                                                        [3,3]        
    [2,2,2]     [2,2,3]             [2,3,3]                                                      [3,3,3]
</pre>
Look at the position of [2]<br/>
Because we choose 2, we can still choose 2, 3<br/>

Look at the position of [2, 3]<br/>
Because we choose 3, we can only choose 3<br/>
And can't choose 2 again<br/>

Look at the position of [3]<br/>
Because we choose 3, we can only choose 3<br/>
And can't choose 2 again<br/>

# Complexity Analysis

n = the legnth of candidates array<br/>
t = the target<br/>

## Time Complexity: O(n ^ t)
- Branching Factor = n
  - At worst, we can decompose from index 0 to n - 1
- Depth = t
  - At worst, we can decompose t times (if the smallest element is 1)
- Each call stack = O(n)

## Space Complexity: O(t)