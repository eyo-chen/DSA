# Problem Explanation

There are two solutions to this problem<br/>

## Solution 1: Two brancing factors(choose and not choose)

At each call stack, we have two choices (branching factor)<br/>
1. Choose the current element
2. Not choose the current element

### Recursive Tree Visualization
Suppose k = 3, n = 7<br/>
(n) means the remaining value<br/>
x means not choose<br/>
<pre>
                                                        7
                                    x(7)                               1(6)
                    x(7)                      2(5)              x(6)                   2(4)
             x(7)        3(4)         x(5)       3(2)     x(6)        3(3)        x(4)        3(1)   
                                                                                  4(0)
</pre>
Only the path 7 -> 1(6) -> 2(4) -> x(4) -> 4(0) is a valid path<br/>
result = [1,2,4]<br/>

At level one, we choose 1<br/>
At level two, we choose 2<br/>
....


### Complexity Analysis

n = input n<br/>
k = input k

#### Time Complexity: O(2 ^ 9) = O(1)
- Branching Factor = 2
  - At each level, we have two choices
- Depth = 9
  - At worst, we go down to level 9
- Each call stack = O(n)

#### Space Complexity: O(1)

## Solution 2: One branching factor (choose)

The idea is for each call stack, we choose from n to 9<br/>
Try all the possiblities<br/>

### Recursive Tree Visualization
<pre>
                                                        7
                           1      2      3      4      5      6      7      8      9
                          2..    3..    4..    5..    6..    7..    8..    9..   
</pre>

### Complexity Analysis

n = input n<br/>
k = input k

#### Time Complexity: O(9 ^ 9) = O(1)
- Branching Factor = 9
  - At each level, we have at most nine choices
- Depth = 9
  - At worst, we go down to level 9
- Each call stack = O(n)

#### Space Complexity: O(1)
