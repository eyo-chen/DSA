# Problem Explanation

## Use Array
The idea is to store the values of the linked list in an array and then find the maximum sum of pairs.<br>
Once we have the array, it's easy to find the maximum sum of pairs<br>

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(n)

## Use Stack
The idea is to use a stack to store the first half of the linked list and then find the maximum sum of pairs.<br>
After we have the stack, we can iterate over the second half of the linked list and pop the stack to find the maximum sum of pairs.<br>

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(n/2) = O(n)

## Reverse the Second Half
The idea is to reverse the second half of the linked list and then find the maximum sum of pairs.<br>
After we have the reversed second half, we can iterate over the first half and the reversed second half to find the maximum sum of pairs.<br>

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(1)
