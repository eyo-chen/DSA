# Problem Explanation

This solution is pretty straightforward<br>
Because digits are stored in reverse order, so we can just simply add each digit number from the head to tail<br>

The small edge case of the problem is caring about the carry when while-loop is breaking<br>
For example,<br>
l1: 1 -> 3 -> 6<br>
l2: 2 -> 5 -> 5<br>
result: 6 -> 8 -> 1 -> 1<br>
As we can see, after while-loop, we only have three digits in the result because we won't have a chance to add finial digit<br>
So we have to check if curCarry is greater than 0, which means the last digits create curCarry<br>
so we have to add it<br>

## Complexity Analysis

### Time Complexity: O(n)

### Space Complexity: O(1)