# Problem Explanation

This solution is pretty straightforward
Because digits are stored in reverse order, so we can just simply add each digit number from the head to tail

The small edge case of the problem is caring about the carry when while-loop is breaking
For example,
l1: 1 -> 3 -> 6
l2: 2 -> 5 -> 5
result: 6 -> 8 -> 1 -> 1
As we can see, after while-loop, we only have three digits in the result because we won't have a chance to add finial digit
So we have to check if curCarry is greater than 0, which means the last digits create curCarry
so we have to add it

## Complexity Analysis

### Time Complexity: O(n)

### Space Complexity: O(1)